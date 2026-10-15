"""Independent Python reference for docs/HASHING.md. Generates expected values
for conformance/hashing/cases.json from case definitions, so the Go
implementation is checked against a second implementation, not itself."""
import hashlib, json, sys, unicodedata

WS = set([*range(0x09, 0x0E), 0x20, 0x85, 0xA0, 0x1680, *range(0x2000, 0x200B),
          0x2028, 0x2029, 0x202F, 0x205F, 0x3000])

def norm(s):
    s = unicodedata.normalize("NFC", s)
    out, inws = [], False
    for ch in s:
        if ord(ch) in WS:
            if not inws:
                out.append(" ")
            inws = True
        else:
            out.append(ch); inws = False
    return "".join(out).strip(" ")

def enc_str(s):
    b = ['"']
    for ch in s:
        if ch == '"': b.append('\\"')
        elif ch == '\\': b.append('\\\\')
        elif ord(ch) < 0x20: b.append('\\u%04x' % ord(ch))
        else: b.append(ch)
    b.append('"')
    return "".join(b)

def enc(v):
    if v is None: return "null"
    if isinstance(v, str): return enc_str(v)
    if isinstance(v, list): return "[" + ",".join(enc(x) for x in v) + "]"
    if isinstance(v, dict):
        return "{" + ",".join(enc_str(k) + ":" + enc(v[k]) for k in sorted(v, key=lambda k: k.encode())) + "}"
    raise TypeError(v)

def sha(b): return "sha256:" + hashlib.sha256(b).hexdigest()

def sval(r, k):
    v = r.get(k)
    if v is None: return ""
    if not isinstance(v, str): raise TypeError(k)
    return norm(v)

def lval(r, k):
    v = r.get(k)
    if v is None: return []
    if not isinstance(v, list): raise TypeError(k)
    return ["" if x is None else norm(x) for x in v]

def gen_obj(r):
    return {"intent": sval(r, "intent"), "constraints": lval(r, "constraints"),
            "affected_scope": lval(r, "affected_scope"), "forbidden_scope": lval(r, "forbidden_scope")}

def spec_hash(content):
    return sha(content.encode().replace(b"\r\n", b"\n"))

def content_obj(r, files):
    t = sval(r, "type") or "blueprint"
    specs = []
    for s in r.get("associated_specs") or []:
        p = s.get("path") if isinstance(s, dict) else None
        specs.append(spec_hash(files[p]) if p is not None and p in files else None)
    return {"type": t, "intent": sval(r, "intent"),
            "constraints": lval(r, "constraints"), "affected_scope": lval(r, "affected_scope"),
            "forbidden_scope": lval(r, "forbidden_scope"), "implements": sval(r, "implements"),
            "associated_specs": specs}

def fill(case):
    k = case["kind"]
    if k == "normalize":
        case["output"] = norm(case["input"])
    elif k == "gen_hash":
        c = enc(gen_obj(case["record"])); case["canonical"] = c; case["hash"] = sha(c.encode())
    elif k == "content_hash":
        c = enc(content_obj(case["record"], case.get("spec_files", {})))
        case["canonical"] = c; case["hash"] = sha(c.encode())
    elif k == "spec_hash":
        case["hash"] = spec_hash(case["content"])
    return case

if __name__ == "__main__":
    src = json.load(open(sys.argv[1]))
    src["cases"] = [fill(c) for c in src["cases"]]
    json.dump(src, open(sys.argv[2], "w"), ensure_ascii=False, indent=2)
    open(sys.argv[2], "a").write("\n")
