import sys
import json
D="prov-2026-20c3360e"
C=[]
def add(id,kind,why,**kw): C.append({"id":id,"kind":kind,"decision":D,"why":why,**kw})
# normalize
N=lambda id,why,s: add("normalize/"+id,"normalize",why,input=s)
N("trim-and-collapse","Whitespace carries no meaning in a record's text, so re-wrapping or re-indenting must not look like a content change.","  Users  can\tsign\n\n in.  ")
N("only-whitespace-is-empty","A string of only whitespace normalizes to the empty string, the same value as missing.","\n \t ")
N("nfc-combining-mark","Two editors can type the same accented letter as one or two code points; NFC makes them the same bytes.","Café")
N("nfc-singleton","NFC also rewrites singletons such as U+212B ANGSTROM SIGN to U+00C5. An implementation that only composes combining marks misses it.","10 Å")
N("nfc-is-not-nfkc","Compatibility characters stay as written. NFKC would fold the ligature to 'fi' and the full-width digit to '1'.","ﬁle １")
N("nbsp-is-whitespace","U+00A0 is White_Space and pasted text often carries it; it collapses like a space.","a  b")
N("unicode-spaces-are-whitespace","Every White_Space code point counts: NEL, ogham, en quad through hair space, line and paragraph separators, narrow no-break, medium mathematical, ideographic.","a\u0085b c d e f g h i　j\u000bk\u000cl")
N("zero-width-and-bom-are-not-whitespace","U+200B, U+FEFF, U+180E and U+2060 are not White_Space. JavaScript's \\s matches U+FEFF; an implementation using it would strip this BOM and disagree.","﻿a​b᠎c⁠d")
N("case-and-punctuation-kept","Only whitespace is normalized. Case, punctuation and quoting are content.","  Don't  USE \"eval\"! ")
# gen_hash
G=lambda id,why,r: add("gen/"+id,"gen_hash",why,record=r)
G("example-from-spec","The worked example in HASHING.md section 5.",{"intent":"Users can\nsign in.\n","constraints":["No  passwords "],"affected_scope":None,"forbidden_scope":None})
G("all-missing","A record with none of the four fields still hashes: every key is present, empty, never omitted.",{})
G("null-equals-missing",  "Null, missing and empty are one value. These three cases share one hash.",{"intent":None,"constraints":None,"affected_scope":None,"forbidden_scope":None})
G("empty-equals-missing", "Null, missing and empty are one value. These three cases share one hash.",{"intent":"","constraints":[],"affected_scope":[],"forbidden_scope":[]})
G("scope-order-ab","List order is kept: sedum reads lists in order, so a reorder is a change. Compare scope-order-ba.",{"intent":"x","affected_scope":["a.go","b.go"]})
G("scope-order-ba","List order is kept: sedum reads lists in order, so a reorder is a change. Compare scope-order-ab.",{"intent":"x","affected_scope":["b.go","a.go"]})
G("duplicates-and-empties-kept","Items are never deduplicated or dropped. An empty or null item stays in place as \"\".",{"constraints":["a","a","",None,"  "]})
G("ignores-non-gen-fields","gen_hash reads only its four fields. Status, version, title, unknown keys and everything else must not move it. Same hash as example-from-spec.",{"intent":"Users can sign in.","constraints":["No passwords"],"status":"implemented","version":7,"reason":"r","title":"T","implements_version":3,"x_unknown":{"a":1}})
G("block-scalar-equals-one-line","A folded or literal block scalar decodes with newlines; after normalization it hashes like the one-line form. Same hash as example-from-spec.",{"intent":"  Users can\n  sign in.\n\n","constraints":["No passwords"]})
G("escapes","Only quote, backslash and U+0000-U+001F are escaped, and always as lowercase \\u00xx. U+0008 is \\u0008, never \\b. '/', '<', '>', '&' and non-ASCII are raw.",{"intent":"say \"hi\" \\ </script> & café 日本 \u0008\u001f\u0001","constraints":[" collapsed "]})
G("all-four-fields","A representative imprint.",{"intent":"Add the users index action.","constraints":["Return JSON","Paginate at 50"],"affected_scope":["app/controllers/users_controller.rb","config/routes.rb"],"forbidden_scope":["db/**"]})
# content_hash
K=lambda id,why,r,files=None: add("content/"+id,"content_hash",why,record=r,**({"spec_files":files} if files is not None else {}))
base={"title":"Users index","type":"imprint","intent":"Add the users index action.","constraints":["Return JSON"],"affected_scope":["app/users.rb"],"forbidden_scope":[],"implements":"prov-2026-1a2b3c4d","implements_version":1,"related":["prov-2026-00000001"],"tags":["api"],"associated_specs":[{"path":"specs/users.linespec","type":"linespec"}]}
F={"specs/users.linespec":"TEST users\nEXPECT 200\n"}
K("all-missing","All five keys are present even for an empty record.",{})
K("baseline","A representative imprint with one spec.",base,F)
import copy
ex=copy.deepcopy(base); ex.update({"id":"prov-2026-ffffffff","status":"implemented","sealed_at_sha":"abc","version":4,"reason":"because","implements_version":2,"author":"a@b","created_at":"2026-01-01","aliases":["prov-2025-aaaaaaaa"],"locked":True,"monitors":["m"],"associated_traces":["t"],"supersedes":"prov-2025-bbbbbbbb","superseded_by":"","extends":"x","x_unknown":[1,2]})
ex["associated_specs"]=[{"path":"specs/users.linespec","type":"rspec","run_command":"rm -rf /","description":"d"}]
K("excluded-fields-ignored","Lifecycle, identity, labels, metadata, migration, governance, attachment, v1-only and unknown fields, plus spec type/run_command/description, are outside content_hash. Same hash as baseline.",ex,F)
K("spec-crlf-equals-lf","Spec bytes are hashed with CRLF folded to LF, so a checkout with autocrlf does not change the record. Same hash as baseline.",base,{"specs/users.linespec":"TEST users\r\nEXPECT 200\r\n"})
mv=copy.deepcopy(base); mv["associated_specs"]=[{"path":"linespecs/moved.linespec"}]
K("spec-path-not-hashed","Moving a spec without changing it is not a content change: paths mean nothing once published. Same hash as baseline.",mv,{"linespecs/moved.linespec":F["specs/users.linespec"]})
K("spec-edit-changes-hash","Editing an integration spec versions its owning record (DESIGN.md, Associated specs).",base,{"specs/users.linespec":"TEST users\nEXPECT 201\n"})
K("spec-missing-is-null","An unreadable spec contributes null, so the record still hashes; creating the file later changes the hash.",base,{})
np=copy.deepcopy(base); np["associated_specs"]=[{"type":"linespec"},{"path":"specs/users.linespec"}]
K("spec-without-path-is-null","An entry with no path contributes null in its position; order of the rest is kept.",np,F)
tg=copy.deepcopy(base); tg.update({"title":"Renamed","tags":["api","v2"],"related":[]})
K("labels-are-not-content","Only behavioral fields are content (Caleb, #provisio 58861). Renaming, retagging or re-linking a record forces no version. Same hash as baseline.",tg,F)
lk=copy.deepcopy(base); lk.update({"type":"blueprint","implements":"prov-2026-99999999","implements_repo":"github.com/x/y"})
K("type-and-parent-are-not-content","type, implements and implements_repo are lineage and classification, not behaviour (provisio-mgr 58864). Same hash as baseline.",lk,F)
cs=copy.deepcopy(base); cs["constraints"]=["Return JSON","Paginate at 50"]
K("constraint-is-content","A constraint changes behaviour, so it changes content_hash.",cs,F)
# spec_hash
S=lambda id,why,c: add("spec/"+id,"spec_hash",why,content=c)
S("empty","An empty file has a hash; it is not the same as a missing file (null).","")
S("lf","Plain LF content.","a\nb\n")
S("crlf-equals-lf","CRLF pairs fold to LF. Same hash as spec/lf.","a\r\nb\r\n")
CH=lambda id,why,regions: C.append({"id":"code/"+id,"kind":"code_hash","decision":"prov-2026-7752337c","why":why,"regions":regions})
CH("no-regions","A record with no regions still has a code hash: the empty list.",[])
CH("order-independent-ab","Region order in files does not matter. Same hash as order-independent-ba.",[{"gen":"sha256:aa","pkg":"sha256:01"},{"gen":"sha256:bb","pkg":"sha256:02"}])
CH("order-independent-ba","Region order in files does not matter. Same hash as order-independent-ab.",[{"gen":"sha256:bb","pkg":"sha256:02"},{"gen":"sha256:aa","pkg":"sha256:01"}])
CH("seeded-regions-ignored","Seeded regions are never regenerated, so they do not count, even unstamped ones. Same hash as order-independent-ab.",[{"gen":"sha256:aa","pkg":"sha256:01","tier":"owned"},{"gen":"sha256:bb","pkg":"sha256:02"},{"gen":"","pkg":"","tier":"seeded"}])
CH("duplicates-kept","Two regions with the same gen and pkg are still two regions.",[{"gen":"sha256:aa","pkg":"sha256:01"},{"gen":"sha256:aa","pkg":"sha256:01"}])
CH("missing-pkg-is-unhashable","A region written before pkg existed has no code hash: absent means stale, and evidence reports the code as unknown (provisio-mgr). Run sedum grow --stamp.",[{"gen":"sha256:aa"},{"gen":"sha256:bb","pkg":"sha256:02"}])
CH("empty-gen-is-unhashable","An empty gen is as unknown as a missing one.",[{"gen":"","pkg":"sha256:01"}])
CH("pkg-change-changes-hash","A template change re-renders code, so the evidence must go stale. Differs from order-independent-ab only in one pkg.",[{"gen":"sha256:aa","pkg":"sha256:09"},{"gen":"sha256:bb","pkg":"sha256:02"}])
S("lone-cr-kept","Only CRLF pairs fold. A lone CR stays, and so does a BOM.","﻿a\rb\n")

# decode (hand-written expectations; no generator)
def Dc(id,why,yaml,fields=None,err=None):
    kw={"yaml":yaml}
    if err: kw.update({"error":True,"error_names":err})
    else: kw["fields"]=fields
    add("decode/"+id,"decode",why,**kw)
Dc("plain-scalars-keep-source-text","Scalars in string fields are their source text, never a typed value. A YAML library that resolves 1.50 to a float and prints 1.5, or 007 to 7, or a date to a timestamp, hashes a different record.",
   "id: prov-2026-00000001\nintent: 1.50\nconstraints:\n  - 1.50\n  - 007\n  - true\n  - 2026-10-15\n  - 0x1F\n  - .inf\n",
   {"intent":"1.50","constraints":["1.50","007","true","2026-10-15","0x1F",".inf"]})
Dc("null-forms","Plain null, Null, NULL, ~ and an empty value are null; a quoted 'null' is the four-letter string. Null then hashes like missing.",
   "intent: ~\nimplements:\nconstraints:\n  - null\n  - Null\n  - NULL\n  - ~\n  - \"null\"\n  - ''\naffected_scope:\nforbidden_scope: null\n",
   {"intent":None,"implements":None,"constraints":[None,None,None,None,"null",""],"affected_scope":None,"forbidden_scope":None})
Dc("block-scalars","Block scalars decode as YAML defines (literal keeps newlines, folded joins lines, chomping indicators apply). Normalization later makes all three hash alike.",
   "intent: |\n  Users can\n  sign in.\nconstraints:\n  - >\n    folded\n    line\n  - |-\n    stripped\n",
   {"intent":"Users can\nsign in.\n","constraints":["folded line\n","stripped"]})
Dc("quoted-scalars","Quoted scalars use their unescaped content.",
   "intent: 'it''s'\nconstraints:\n  - \"tab\\there\"\n  - \"caf\\u00e9\"\n",
   {"intent":"it's","constraints":["tab\there","caf\u00e9"]})
Dc("anchors-and-comments","Aliases resolve as YAML defines; comments never matter.",
   "# header comment\nintent: &i Shared text # trailing\nconstraints:\n  - *i\n",
   {"intent":"Shared text","constraints":["Shared text"]})
Dc("keys-are-case-sensitive","Key names match exactly. affectedScope is an unknown key: preserved, never hashed.",
   "affectedScope:\n  - a.go\naffected_scope:\n  - b.go\n",
   {"affected_scope":["b.go"]})
Dc("duplicate-key-is-error","A duplicate top-level key would make the hash depend on which one a parser keeps.",
   "intent: a\nintent: b\n",err="duplicate")
Dc("scalar-for-list-is-error","A scalar where a list is expected is an error, not a one-item list.",
   "affected_scope: app/users.rb\n",err="affected_scope")
Dc("mapping-for-string-is-error","A mapping where a string is expected is an error.",
   "intent:\n  text: hi\n",err="intent")
json.dump({"format":"provenance.hashing","corpus_version":3,"description":"Golden cases for canonical record hashing (docs/HASHING.md): string normalization, gen_hash, content_hash, spec hashes and YAML decoding. Every case cites the record that decided it.","cases":C},open(sys.argv[1],"w"),ensure_ascii=False,indent=2)
