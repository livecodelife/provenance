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

# canonical (bool support, prov-2026-32ee2643). Not a record hash: the shared
# encoder for other objects, such as sedum's pkg hash.
D2="prov-2026-32ee2643"
def Cn(id,why,value,err=False):
    kw={"value":value}
    if err: kw["error"]=True
    C.append({"id":"canonical/"+id,"kind":"canonical","decision":D2,"why":why,**kw})
Cn("bools-in-map","A boolean is the bare literal true or false, never a quoted string.",{"required":True,"default":False})
Cn("bool-in-list","Booleans inside a list keep their position and sit beside strings and null.",[True,False,"true",None])
Cn("keys-sorted-around-bools","Key order is by UTF-8 bytes regardless of value type.",{"b":True,"a":False,"B":True,"aa":None,"":False})
Cn("bare-bool","A top-level boolean encodes alone.",True)
Cn("string-true-stays-string","The string true is still a string: types are not coerced in either direction.",{"s":"true","b":True})
Cn("nested","Booleans nest through maps and lists like any other value.",{"kwargs":{"x":{"type":"string","required":True,"default":None},"y":{"type":"bool","required":False,"default":False}},"list":[[True],[]]})
Cn("number-rejected","A float is still not canonical and is an error, never silently stringified. (Integers arrived in corpus version 5; see canonical_int.)",{"n":1.5},err=True)

# canonical_int (integer support, prov-2026-cb479524). An integer is carried as
# {"$int": text} so no language rounds a 64-bit value through a float.
D3="prov-2026-cb479524"
def Ci(id,why,value,err=False):
    kw={"value":value}
    if err: kw["error"]=True
    C.append({"id":"canonical_int/"+id,"kind":"canonical_int","decision":D3,"why":why,**kw})
I=lambda t:{"$int":t}
Ci("ints-in-map","An integer is bare base-10 digits; zero is 0 and a negative is -N.",{"n":I("3"),"neg":I("-3"),"zero":I("0")})
Ci("int64-bounds","The full int64 range encodes exactly.",[I("-9223372036854775808"),I("9223372036854775807")])
Ci("uint64-max","The full uint64 range encodes exactly, which no float64 can.",I("18446744073709551615"))
Ci("beyond-2-53","2^53+1 is not representable as a double; an encoder that goes through a float writes the wrong digits.",I("9007199254740993"))
Ci("negative-zero-normalised","-0 is written 0.",I("-0"))
Ci("plus-and-leading-zeros-normalised","Input text may carry a plus sign or leading zeros; the encoding never does.",[I("+5"),I("007")])
Ci("mixed-with-bool-and-string","Integers sit beside booleans, strings and null; a string of digits stays a string.",{"default":I("3"),"required":True,"s":"3","type":"int","none":None})
Ci("float-text-rejected","Only integers were added: 1.5 is an error.",{"n":I("1.5")},err=True)
Ci("exponent-text-rejected","1e3 is a float literal, not an integer: an error.",I("1e3"),err=True)
Ci("uint64-overflow-rejected","One past uint64 max is out of range: an error, never a wrapped or rounded value.",I("18446744073709551616"),err=True)
Ci("int64-underflow-rejected","One below int64 min is out of range: an error.",I("-9223372036854775809"),err=True)
Ci("empty-text-rejected","No digits is not an integer.",I(""),err=True)
# record-file specs (prov-2026-4234f0bf)
D4="prov-2026-4234f0bf"
def R(id,why,r,files):
    C.append({"id":"content/"+id,"kind":"content_hash","decision":D4,"why":why,"record":r,"spec_files":files})
rf=copy.deepcopy(base); rf["associated_specs"]=[{"path":"provenance/prov-2026-0000b001.yml"}]
R("record-file-spec-other-record","A spec that is another record's file contributes the constant marker record-file/v1, not the file bytes: rewriting the record store never changes the hash.",rf,{"provenance/prov-2026-0000b001.yml":"version: 3\n"})
R("record-file-spec-bytes-ignored","Same hash as record-file-spec-other-record with different file bytes.",rf,{"provenance/prov-2026-0000b001.yml":"version: 4\ncontent_hash: x\n"})
R("record-file-spec-missing-is-marker","The file need not exist: a record-file spec is the marker, never null. Same hash as record-file-spec-other-record.",rf,{})
sf=copy.deepcopy(base); sf["id"]="prov-2026-0000b002"; sf["associated_specs"]=[{"path":"provenance/prov-2026-0000b002.yml"}]
R("record-file-spec-self","A record that names its own file as a spec hashes it as the marker too (a record cannot hash its own bytes).",sf,{"provenance/prov-2026-0000b002.yml":"id: prov-2026-0000b002\n"})
dr=copy.deepcopy(base); dr["associated_specs"]=[{"path":"other/dir/prov-2026-0000b001-api.yml"}]
R("record-file-spec-any-directory-and-suffix","Detection is by base name alone: any directory, an optional -suffix. Same hash as record-file-spec-other-record.",dr,{})
ll=copy.deepcopy(base); ll["associated_specs"]=[{"path":"specs/prov-2026-0000b001.txt"},{"path":"specs/my-prov-2026-0000b001.yml"}]
R("lookalike-names-are-ordinary-specs","Only a base name of exactly prov-YYYY-<hex8>[-suffix].yml is a record file; anything else is hashed by its bytes.",ll,{"specs/prov-2026-0000b001.txt":"a\n","specs/my-prov-2026-0000b001.yml":"b\n"})
mx=copy.deepcopy(base); mx["associated_specs"]=[{"path":"specs/users.linespec"},{"path":"provenance/prov-2026-0000b001.yml"},{"type":"linespec"}]
R("record-file-spec-keeps-its-slot","The marker sits in the spec's own list position, between an ordinary spec and a pathless entry (null).",mx,F)
# published versions (prov-2026-29d43aaf): the lexicon form of a version record
D7="prov-2026-29d43aaf"
def V(id,why,version,stated="match"):
    C.append({"id":"version/"+id,"kind":"version_hash","decision":D7,"why":why,"version":version,"stated":stated})
vc={"title":"Users index","type":"imprint","createdAt":"2026-10-01","author":"dev@example.com","intent":"Add the users index action.",
    "constraints":["Return JSON"],"affectedScope":["app/users.rb"],"forbiddenScope":["app/admin.rb"],"related":["prov-2026-0000c001"],"tags":["api"],
    "aliases":["prov-2026-0000d001"],"implements":"prov-2026-0000b001","implementsRepo":"shared","implementsVersion":1,
    "associatedSpecs":[{"path":"specs/users.linespec","type":"linespec"},{"path":"provenance/prov-2026-0000e001.yml"},{"path":"","type":"linespec"}]}
vh=["sha256:55b89d8c9e1e8062aff77fee7c0464c66c3fecad32899d82dcbff1d4667aa1e1","record-file/v1",""]
def ver(content=None,hashes=None,**kw):
    v={"recordId":"prov-2026-0000a001","version":1,"author":"dev@example.com","createdAt":"2026-10-01T09:00:00Z",
       "specHashes":copy.deepcopy(hashes if hashes is not None else vh),"content":copy.deepcopy(content if content is not None else vc)}
    v.update(kw); return v
V("all-fields","A published version hashes from its content and specHashes alone: the five content keys, with each spec hash in its list slot, the record-file marker as is, and the empty string (no hash) as null.",ver())
lab=copy.deepcopy(vc); lab.update({"title":"Renamed","tags":[],"related":[],"aliases":[],"implements":"prov-2026-0000b002","implementsVersion":7,"author":"other@example.com","locked":True})
V("labels-and-pin-ignored","Labels, cross-references, the parent pin and metadata are published but are not content: same hash as all-fields.",ver(content=lab,reason="Renamed",reasonKind="fix",parentHash="sha256:"+"0"*64))
ti=copy.deepcopy(vc); ti["intent"]="Add the users index action, with paging."
V("tampered-intent","A version whose content no longer matches its stated contentHash is invalid; the hash is the one recomputed from the content.",ver(content=ti),stated="match-of-all-fields")
V("tampered-spec-hash","A spec hash that differs from the one the contentHash was made with is detected the same way.",ver(hashes=["sha256:"+"1"*64,"record-file/v1",""]),stated="match-of-all-fields")
V("spec-count-mismatch","specHashes must have one entry per associated spec: a different count cannot be verified and is an error.",ver(hashes=vh[:2]),stated="match-of-all-fields")
ne=copy.deepcopy(vc); ne["intent"]="  Caf\u00e9  \u2014  na\u00efve\n\n  users  "
V("text-is-normalized","The intent is normalized as in content_hash: whitespace collapsed, NFC, so wrapping differences in a published version do not change the hash.",ver(content=ne))
em={"title":"","intent":""}
V("empty-content","Only the required fields: the same canonical bytes and hash as content/all-missing.",ver(content=em,hashes=[]))
ps=copy.deepcopy(em); ps["associatedSpecs"]=[{"path":"","type":"linespec"}]
V("pathless-spec-is-null","A spec with no path publishes path \"\" and the empty hash; both hash as null in the spec's slot.",ver(content=ps,hashes=[""]))
json.dump({"format":"provenance.hashing","corpus_version":7,"description":"Golden cases for canonical record hashing (docs/HASHING.md): string normalization, gen_hash, content_hash, spec hashes, published version records and YAML decoding. Every case cites the record that decided it.","cases":C},open(sys.argv[1],"w"),ensure_ascii=False,indent=2)
