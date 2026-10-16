"""Independent Python reference for docs/HASHING.md. Generates expected values
for conformance/hashing/cases.json from case definitions, so the Go
implementation is checked against a second implementation, not itself."""
import hashlib, json, re, sys, unicodedata

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

import re
INTS = False  # integers are canonical only in kind canonical_int (markers)
def conv(v):
    """Replace {"$int": text} markers by Python ints; bad text is an error."""
    if isinstance(v, dict):
        if list(v) == ["$int"]:
            t = v["$int"]
            if not re.fullmatch(r"[+-]?[0-9]+", t): raise TypeError(t)
            n = int(t)
            if not -(2**63) <= n <= 2**64 - 1: raise TypeError(t)
            return n
        return {k: conv(x) for k, x in v.items()}
    if isinstance(v, list): return [conv(x) for x in v]
    return v

def enc(v):
    if v is None: return "null"
    if isinstance(v, bool): return "true" if v else "false"
    if isinstance(v, int):
        if not INTS: raise TypeError(v)
        return str(v)
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
    specs = []
    for s in r.get("associated_specs") or []:
        p = s.get("path") if isinstance(s, dict) else None
        if p is not None and re.match(r"^prov-[0-9]{4}-[0-9a-f]{8}(-[A-Za-z0-9._-]+)?\.yml\Z", p.split("/")[-1]):
            specs.append("record-file/v1")
        else:
            specs.append(spec_hash(files[p]) if p is not None and p in files else None)
    return {"intent": sval(r, "intent"),
            "constraints": lval(r, "constraints"), "affected_scope": lval(r, "affected_scope"),
            "forbidden_scope": lval(r, "forbidden_scope"), "associated_specs": specs}

ALL_FIELDS_HASH = [None]

def owned(regions):
    return [r for r in regions if r.get("tier", "owned") == "owned"]

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
    elif k == "canonical":
        try:
            c = enc(case["value"]); case["canonical"] = c; case["hash"] = sha(c.encode())
        except TypeError:
            case["error"] = True
    elif k == "canonical_int":
        global INTS
        INTS = True
        try:
            c = enc(conv(case["value"])); case["canonical"] = c; case["hash"] = sha(c.encode())
        except TypeError:
            case["error"] = True
        finally:
            INTS = False
    elif k == "version_hash":
        v = case["version"]
        content, hashes = v["content"], v["specHashes"]
        if len(hashes) != len(content.get("associatedSpecs") or []):
            case["error"] = True
            del case["stated"]
            return case
        rec = {"intent": content.get("intent"), "constraints": content.get("constraints"),
               "affected_scope": content.get("affectedScope"), "forbidden_scope": content.get("forbiddenScope")}
        obj = gen_obj(rec)
        obj["associated_specs"] = [h if h != "" else None for h in hashes]
        c = enc(obj)
        case["canonical"] = c; case["hash"] = sha(c.encode())
        if case["stated"] == "match":
            v["contentHash"] = case["hash"]
        else:
            v["contentHash"] = ALL_FIELDS_HASH[0]
        case["valid"] = v["contentHash"] == case["hash"]
        del case["stated"]
        if case["id"] == "version/all-fields":
            ALL_FIELDS_HASH[0] = case["hash"]
    elif k == "code_hash" and any(not r.get("gen") or not r.get("pkg") for r in owned(case["regions"])):
        case["error"] = True
    elif k == "code_hash":
        regions = sorted(enc({"gen": r.get("gen", ""), "pkg": r.get("pkg", "")}) for r in owned(case["regions"]))
        c = "[" + ",".join(regions) + "]"
        case["canonical"] = c; case["hash"] = sha(c.encode())
    return case

if __name__ == "__main__":
    src = json.load(open(sys.argv[1]))
    src["cases"] = [fill(c) for c in src["cases"]]
    json.dump(src, open(sys.argv[2], "w"), ensure_ascii=False, indent=2)
    open(sys.argv[2], "a").write("\n")
