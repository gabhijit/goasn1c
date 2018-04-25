## Test Cases from the asn1c

### Summary
| Test Case Type | Pass Percentage |
|----|----:|
| Lexer (tokenization) | 99.25 |
| Parser  | 48.00 |
| Code Generation | 0.0 |

### Legend
| value| meaning | notes |
|------| ---- | --- |
| Y | Passes|
| X | Not Supported (Won't be implemented)|
| ? | Kind of Works |
| - | Not implemented|
| F | Fails Right now |

### Results

| Test Case | Tokenizer Support | Parser Support | Code Generation | Notes |
|:---|:---:|:---:|:---:|:---:|
| 00-empty-OK.asn1 | Y | Y | - |
| 01-empty-OK.asn1 |  Y | Y | |
| 02-garbage-NP.asn1 |  Y | Y | |
| 03-enum-OK.asn1 |  Y | Y | |
| 04-enum-SE.asn1 |  Y | Y | |
| 05-enum-SE.asn1 |  Y | Y | |
| 06-enum-SE.asn1 |  Y | Y | |
| 07-int-OK.asn1 |  Y |  Y | |
| 08-int-SE.asn1 |  Y | Y | |
| 09-int-SE.asn1 |  Y | Y | |
| 10-int-OK.asn1 |  Y | Y | |
| 11-int-SE.asn1 |  Y | Y | |
| 12-int-SE.asn1 |  Y | Y | |
| 13-resolver-OK.asn1 |  Y | Y | |
| 14-resolver-OK.asn1 |  Y | Y | |
| 15-resolver-SE.asn1 |  Y | Y | |
| 16-constraint-OK.asn1 |  Y |  | |
| 17-tags-OK.asn1 |  Y |  Y | |
| 18-class-OK.asn1 |  Y |  | |
| 19-param-OK.asn1 |  Y |  | |
| 20-constr-OK.asn1 |  Y |  | |
| 21-tags-OK.asn1 |  Y | Y | |
| 22-tags-OK.asn1 |  Y |  | |
| 23-bits-OK.asn1 |  Y |  | |
| 24-sequence-OK.asn1 |  Y | Y | |
| 25-misc-OK.asn1 |  Y |  | |
| 26-sequence-SE.asn1 |  Y | Y | |
| 27-set-SE.asn1 |  Y | Y | |
| 28-tags-SE.asn1 |  Y | Y | |
| 29-tags-OK.asn1 |  Y | Y | |
| 30-set-OK.asn1 |  Y | Y | |
| 31-set-of-OK.asn1 |  Y |  | |
| 32-sequence-of-OK.asn1 |  Y |  | |
| 33-misc-OK.asn1 |  Y |  | |
| 34-class-OK.asn1 |  Y |  | |
| 35-set-choice-OK.asn1 |  Y | Y | |
| 36-indirect-choice-SE.asn1 |  Y | Y | |
| 37-indirect-choice-OK.asn1 |  Y | Y | |
| 38-comments-OK.asn1 |  Y | Y | |
| 39-sequence-of-OK.asn1 |  Y |  | |
| 40-int-optional-SE.asn1 |  Y |  | |
| 41-int-optional-OK.asn1 |  Y |  | |
| 42-real-life-OK.asn1 |  Y |  | |
| 43-recursion-OK.asn1 |  Y |  | |
| 44-choice-in-sequence-OK.asn1 |  Y |  | |
| 45-undefined-type-SE.asn1 |  Y | Y | |
| 46-redefine-OK.asn1 |  Y | Y | |
| 47-set-ext-OK.asn1 |  Y | Y | |
| 48-real-life-OK.asn1 |  Y | Y | |
| 49-real-life-OK.asn1 |  Y | Y | |
| 50-constraint-OK.asn1 |  Y |  | |
| 51-constraint-SE.asn1 |  Y |  Y | |
| 52-constraint-SE.asn1 |  Y |  Y | |
| 53-constraint-SE.asn1 |  Y |  Y | |
| 54-constraint-SE.asn1 |  Y |  Y | |
| 55-components-of-OK.asn1 |  Y |  | |
| 56-components-of-SE.asn1 |  Y |  | |
| 57-components-of-OK.asn1 |  Y |  | |
| 58-param-OK.asn1 |  Y |  | |
| 59-choice-extended-OK.asn1 |  Y | Y | |
| 60-any-OK.asn1 |  Y | Y | |
| 61-any-1-SE.asn1 |  Y | Y | |
| 62-any-OK.asn1 |  Y |  | |
| 63-any-2-SE.asn1 |  Y | Y | |
| 64-oid-constr-OK.asn1 |  Y |  | |
| 65-multi-tag-OK.asn1 |  Y |  | |
| 66-ref-simple-OK.asn1 |  Y |  | |
| 67-embedded-choice-OK.asn1 |  Y | Y | |
| 68-enum-default-OK.asn1 |  Y | Y | |
| 69-reserved-words-OK.asn1 |  Y |  | |
| 70-xer-test-OK.asn1 |  Y |  | |
| 71-duplicate-types-SE.asn1 |  Y | Y | |
| 72-same-names-OK.asn1 |  Y |  | |
| 73-circular-OK.asn1 |  Y |  | |
| 74-int-enum-constraints-OK.asn1 |  Y |  Y | |
| 75-duplicate-modules-SE.asn1 |  Y |  Y | |
| 76-duplicate-modules-SW.asn1 |  Y |  Y | |
| 77-str-default-OK.asn1 |  Y |  | |
| 78-str-default-SE.asn1 |  Y |  | |
| 79-constrained-by-OK.asn1 |  Y |  | |
| 80-chardefs-OK.asn1 |  Y |  | |
| 81-type-default-OK.asn1 |  Y |  | |
| 82-with-comps-OK.asn1 |  Y |  | |
| 83-with-comps-OK.asn1 |  Y |  | |
| 84-param-tags-OK.asn1 |  Y |  | |
| 85-comments-OK.asn1 |  F| | |
| 86-atags-OK.asn1 |  Y | Y | |
| 87-old-syntax-OK.asn1 |  Y | Y | |
| 88-integer-enum-OK.asn1 |  Y | Y | |
| 89-bit-string-enum-OK.asn1 |  Y |  | |
| 90-cond-int-type-OK.asn1 |  Y |  Y | |
| 91-cond-int-blessSize-OK.asn1 |  Y |  | |
| 92-circular-loops-OK.asn1 |  Y |  | |
| 93-asn1c-controls-OK.asn1 |  Y |  | |
| 94-set-optionals-OK.asn1 |  Y |  | |
| 95-choice-per-order-OK.asn1 |  Y | Y | |
| 96-type-identifier-OK.asn1 |  Y |  | |
| 97-type-identifier-SW.asn1 |  Y |  | |
| 98-attribute-class-OK.asn1 |  Y |  | |
| 99-class-sample-OK.asn1 |  Y |  | |
| 100-class-ref-OK.asn1 |  Y |  | |
| 101-class-ref-SE.asn1 |  Y |  | |
| 102-class-ref-SE.asn1 |  Y |  | |
| 103-reference-SE.asn1 |  Y |  | |
| 104-param-1-OK.asn1 |  Y |  | |
| 105-param-2-OK.asn1 |  Y |  | |
| 106-param-constr-OK.asn1 |  Y |  | |
| 107-param-constr-2-OK.asn1 |  Y |  | |
| 108-param-constr-3-OK.asn1 |  Y |  | |
| 109-bit-string-SE.asn1 |  Y |  | |
| 110-param-3-OK.asn1 |  Y |  | |
| 111-param-4-SE.asn1 |  Y |  | |
| 112-param-class-OK.asn1 |  Y |  | |
| 113-bit-string-SE.asn1 |  Y |  | |
| 114-bit-string-SE.asn1 |  Y |  | |
| 115-bit-string-OK.asn1 |  Y |  | |
| 116-bit-string-SE.asn1 |  Y |  | |
| 117-real-constraint-OK.asn1 |  Y |  | |
| 118-per-constraint-OK.asn1 |  Y | Y | |
| 119-per-strings-OK.asn1 |  Y |  | |
| 121-empty-imports-OK.asn1 |  Y |  Y | |
| 122-pattern-OK.asn1 |  Y |  | |
| 123-valueassignment-OK.asn1 |  Y | Y | |
| 124-multiconstraint-OK.asn1 |  Y |  | |
| 125-bitstring-constraint-OK.asn1 |  Y |  | |
| 126-per-extensions-OK.asn1 |  Y | Y | |
| 127-per-long-OK.asn1 |  Y | Y | |
| 128-enum-SE.asn1 |  Y | Y | |
| 129-enum-OK.asn1 |  Y | Y | |
| 129-enum-SE.asn1 |  Y | Y | |
| 130-enum-OK.asn1 |  Y | Y | |
| 131-per-empty-OK.asn1 |  Y | Y | |
| 132-per-choice-OK.asn1 |  Y |  | |
| 133-per-constraints-OK.asn1 |  Y |  | |
