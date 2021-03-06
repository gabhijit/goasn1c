%%
/* grammar from itu-t X-680-2002-07 */
ModuleDefinition: ModuleIdentifier
		DEFINITIONS
		TagDefault
		ExtensionDefault
		" ::= "
		BEGIN
		ModuleBody
		END

ModuleIdentifier: modulereference DefinitiveIdentifier
DefinitiveIdentifier ::= " { " DefinitiveObjIdComponentList " } "
		     | empty
DefinitiveObjIdComponentList ::=
		DefinitiveObjIdComponent
		| DefinitiveObjIdComponent DefinitiveObjIdComponentList

DefinitiveObjIdComponent ::= NameForm | DefinitiveNumberForm | DefinitiveNameAndNumberForm

DefinitiveNumberForm ::= number

DefinitiveNameAndNumberForm ::= identifier " ( " DefinitiveNumberForm " ) "

TagDefault ::= EXPLICIT TAGS | IMPLICIT TAGS | AUTOMATIC TAGS | empty

ExtensionDefault ::= EXTENSIBILITY IMPLIED | empty

/* completed so far */

ModuleBody ::= Exports Imports AssignmentList | empty

Exports ::= EXPORTS SymbolsExported ";" | EXPORTS ALL ";" | empty

SymbolsExported ::= SymbolList | empty

Imports ::= IMPORTS SymbolsImported ";" | empty

SymbolsImported ::= SymbolsFromModuleList | empty

SymbolsFromModuleList ::= SymbolsFromModule | SymbolsFromModuleList SymbolsFromModule

SymbolsFromModule ::= SymbolList FROM GlobalModuleReference

GlobalModuleReference ::= modulereference AssignedIdentifier

AssignedIdentifier ::= ObjectIdentifierValue | DefinedValue | empty

SymbolList ::= Symbol | SymbolList "," Symbol

Symbol ::= Reference | ParameterizedReference

Reference ::= typereference | valuereference | objectclassreference | objectreference | objectsetreference

AssignmentList ::= Assignment | AssignmentList Assignment

Assignment ::= 	TypeAssignment
		| ValueAssignment
		| XMLValueAssignment
		| ValueSetTypeAssignment
		| ObjectClassAssignment
		| ObjectAssignment
		| ObjectSetAssignment
		| ParameterizedAssignment

DefinedType ::=
		ExternalTypeReference
		| typereference
		| ParameterizedType
		| ParameterizedValueSetType

ExternalTypeReference ::= modulereference " . " typereference

NonParameterizedTypeName ::=
		ExternalTypeReference
		| typereference
		| xmlasn1typename

DefinedValue ::=
		ExternalValueReference
		| valuereference
		| ParameterizedValue

ExternalValueReference ::= modulereference " . " valuereference

AbsoluteReference ::= " @ " ModuleIdentifier " . " ItemSpec

ItemSpec ::= typereference | ItemId " . " ComponentId

ItemId ::= ItemSpec

ComponentId ::= identifier | number | " * "

TypeAssignment ::= typereference " ::= " Type

ValueAssignment ::= valuereference Type " ::= " Value

XMLValueAssignment ::= valuereference " ::= " XMLTypedValue

XMLTypedValue ::=
		"<" & NonParameterizedTypeName ">" XMLValue "</" & NonParameterizedTypeName ">"
		| "<" & NonParameterizedTypeName "/>"

ValueSetTypeAssignment ::= typereference Type " ::= " ValueSet

ValueSet ::= " { " ElementSetSpecs " } "

Type ::= BuiltinType | ReferencedType | ConstrainedType

BuiltinType ::=
		BitStringType
		| BooleanType
		| CharacterStringType
		| ChoiceType
		| EmbeddedPDVType
		| EnumeratedType
		| ExternalType
		| InstanceOfType
		| IntegerType
		| NullType
		| ObjectClassFieldType
		| ObjectIdentifierType
		| OctetStringType
		| RealType
		| RelativeOIDType
		| SequenceType
		| SequenceOfType
		| SetType
		| SetOfType
		| TaggedType

NamedType ::= identifier Type

ReferencedType ::=
		DefinedType
		| UsefulType
		| SelectionType
		| TypeFromObject
		| ValueSetFromObjects

Value ::=
		BuiltinValue
		| ReferencedValue
		| ObjectClassFieldValue

XMLValue ::= XMLBuiltinValue | XMLObjectClassFieldValue

BuiltinValue ::=
		BitStringValue
		| BooleanValue
		| CharacterStringValue
		| ChoiceValue
		| EmbeddedPDVValue
		| EnumeratedValue
		| ExternalValue
		| InstanceOfValue
		| IntegerValue
		| NullValue
		| ObjectIdentifierValue
		| OctetStringValue
		| RealValue
		| RelativeOIDValue
		| SequenceValue
		| SequenceOfValue
		| SetValue
		| SetOfValue
		| TaggedValue

XMLBuiltinValue ::=
		XMLBitStringValue
		| XMLBooleanValue
		| XMLCharacterStringValue
		| XMLChoiceValue
		| XMLEmbeddedPDVValue
		| XMLEnumeratedValue
		| XMLExternalValue
		| XMLInstanceOfValue
		| XMLIntegerValue
		| XMLNullValue
		| XMLObjectIdentifierValue
		| XMLOctetStringValue
		| XMLRealValue
		| XMLRelativeOIDValue
		| XMLSequenceValue
		| XMLSequenceOfValue
		| XMLSetValue
		| XMLSetOfValue
		| XMLTaggedValue

ReferencedValue ::= DefinedValue | ValueFromObject

NamedValue ::= identifier Value

XMLNamedValue ::= "<" & identifier ">" XMLValue "</" & identifier ">"

BooleanType ::= BOOLEAN

BooleanValue::= TRUE | FALSE

XMLBooleanValue ::= "<" & "true" "/>" | "<" & "false" "/>"

IntegerType ::= INTEGER | INTEGER " { " NamedNumberList " } "

NamedNumberList ::= NamedNumber | NamedNumberList "," NamedNumber

NamedNumber ::= identifier " ( " SignedNumber " ) " | identifier " ( " DefinedValue " ) "

SignedNumber ::= number | "-" number

IntegerValue ::= SignedNumber | identifier

XMLIntegerValue ::= SignedNumber | "<" & identifier "/>"

EnumeratedType ::= ENUMERATED " { " Enumerations " } "

Enumerations ::=
		RootEnumeration
		| RootEnumeration "," " ... " ExceptionSpec
		| RootEnumeration "," " ... " ExceptionSpec "," AdditionalEnumeration

RootEnumeration ::= Enumeration

AdditionalEnumeration ::= Enumeration

Enumeration ::=
		EnumerationItem
		| EnumerationItem "," Enumeration

EnumerationItem ::= identifier | NamedNumber

EnumeratedValue ::= identifier

XMLEnumeratedValue ::= "<" & identifier "/>"

RealType ::= REAL

RealValue ::= NumericRealValue | SpecialRealValue

NumericRealValue ::=
	realnumber
	| "-" realnumber
	| SequenceValue -- Value of the associated sequence type

SpecialRealValue ::= PLUS-INFINITY | MINUS-INFINITY

XMLRealValue ::= XMLNumericRealValue | XMLSpecialRealValue

XMLNumericRealValue ::= realnumber | "-" realnumber

XMLSpecialRealValue ::= "<" & PLUS-INFINITY "/>" | "<" & MINUS-INFINITY "/>"

BitStringType ::= BIT STRING | BIT STRING " { " NamedBitList " } "

NamedBitList::= NamedBit | NamedBitList "," NamedBit

NamedBit ::= identifier " ( " number " ) " | identifier " ( " DefinedValue " ) "

BitStringValue ::=
		bstring
		| hstring
		| " { " IdentifierList " } "
		| " { " " } "
		| CONTAINING Value

IdentifierList ::= identifier | IdentifierList "," identifier

XMLBitStringValue ::=
		XMLTypedValue
		| xmlbstring
		| XMLIdentifierList
		| empty

XMLIdentifierList ::= "<" & identifier "/>" | XMLIdentifierList "<" & identifier "/>"

OctetStringType ::= OCTET STRING

OctetStringValue ::= bstring | hstring | CONTAINING Value

XMLOctetStringValue ::= XMLTypedValue | xmlhstring

NullType ::= NULL

NullValue ::= NULL

XMLNullValue ::= empty

SequenceType ::=
		SEQUENCE " { " " } "
		| SEQUENCE " { " ExtensionAndException OptionalExtensionMarker " } "
		| SEQUENCE " { " ComponentTypeLists " } "

ExtensionAndException ::= " ... " | " ... " ExceptionSpec

OptionalExtensionMarker ::= "," " ... " | empty

ComponentTypeLists ::=
		RootComponentTypeList
		| RootComponentTypeList "," ExtensionAndException ExtensionAdditions OptionalExtensionMarker
		| RootComponentTypeList "," ExtensionAndException ExtensionAdditions ExtensionEndMarker "," RootComponentTypeList
		| ExtensionAndException ExtensionAdditions ExensionEndMarker "," RootComponentTypeList
		| ExtensionAndException ExtensionAdditions OptionalExtensionMarker

RootComponentTypeList ::= ComponentTypeList

ExtensionEndMarker ::= "," " ... "

ExtensionAdditions ::= "," ExtensionAdditionList | empty

ExtensionAdditionList ::= ExtensionAddition | ExtensionAdditionList "," ExtensionAddition

ExtensionAddition ::= ComponentType | ExtensionAdditionGroup

ExtensionAdditionGroup ::= " [[ " VersionNumber ComponentTypeList " ]] "

VersionNumber ::= empty | number " : "

ComponentTypeList ::= ComponentType | ComponentTypeList "," ComponentType

ComponentType ::= NamedType | NamedType OPTIONAL | NamedType DEFAULT Value | COMPONENTS OF Type

SequenceValue ::= " { " ComponentValueList " } " | " { " " } "

ComponentValueList ::= NamedValue | ComponentValueList "," NamedValue

XMLSequenceValue ::= XMLComponentValueList | empty

XMLComponentValueList ::= XMLNamedValue | XMLComponentValueList XMLNamedValue

SequenceOfType ::= SEQUENCE OF Type | SEQUENCE OF NamedType

SequenceOfValue ::= " { " ValueList " } " | " { " NamedValueList " } " | " { " " } "

ValueList ::= Value | ValueList "," Value

XMLSequenceOfValue ::=
		XMLValueList
		| XMLDelimitedItemList
		| XMLSpaceSeparatedList
		| empty

XMLValueList ::= XMLValueOrEmpty | XMLValueOrEmpty XMLValueList

XMLValueOrEmpty ::= XMLValue | "<" & NonParameterizedTypeName "/>"

XMLSpaceSeparatedList ::= XMLValueOrEmpty | XMLValueOrEmpty " " XMLSpaceSeparatedList

XMLDelimitedItemList ::= XMLDelimitedItem | XMLDelimitedItem XMLDelimitedItemList

XMLDelimitedItem ::=
		"<" & NonParameterizedTypeName ">" XMLValue "</" & NonParameterizedTypeName ">"
		| "<" & identifier ">" XMLValue "</" & identifier ">"

SetType ::=
		SET " { " " } "
		| SET " { " ExtensionAndException OptionalExtensionMarker " } "
		| SET " { " ComponentTypeLists " } "

SetValue ::= " { " ComponentValueList " } " | " { " " } "

XMLSetValue ::= XMLComponentValueList | empty

SetOfType ::= SET OF Type | SET OF NamedType

SetOfValue ::= " { " ValueList " } " | " { " NamedValueList " } " | " { " " } "

XMLSetOfValue ::= XMLValueList | XMLDelimitedItemList | XMLSpaceSeparatedList | empty

ChoiceType ::= CHOICE " { " AlternativeTypeLists " } "

AlternativeTypeLists ::=
		RootAlternativeTypeList
		| RootAlternativeTypeList "," ExtensionAndException ExtensionAdditionAlternatives OptionalExtensionMarker

RootAlternativeTypeList ::= AlternativeTypeList

ExtensionAdditionAlternatives ::= "," ExtensionAdditionAlternativesList | empty

ExtensionAdditionAlternativesList ::=
		ExtensionAdditionAlternative
		| ExtensionAdditionAlternativesList "," ExtensionAdditionAlternative

ExtensionAdditionAlternative ::=
		ExtensionAdditionAlternativesGroup
		| NamedType

ExtensionAdditionAlternativesGroup ::= " [[ " VersionNumber AlternativeTypeList " ]] "

AlternativeTypeList ::= NamedType | AlternativeTypeList "," NamedType

ChoiceValue ::= identifier " : " Value

XMLChoiceValue ::= "<" & identifier ">" XMLValue "</" & identifier ">"

SelectionType ::= identifier "<" Type

TaggedType ::= Tag Type | Tag IMPLICIT Type | Tag EXPLICIT Type

Tag ::= "[" Class ClassNumber "]"

ClassNumber ::= number | DefinedValue

Class ::= UNIVERSAL | APPLICATION | PRIVATE | empty

TaggedValue ::= Value

XMLTaggedValue ::= XMLValue

EmbeddedPDVType ::= EMBEDDED PDV

EmbeddedPDVValue ::= SequenceValue

XMLEmbeddedPDVValue ::= XMLSequenceValue

ExternalType ::= EXTERNAL

ExternalValue ::= SequenceValue

XMLExternalValue ::= XMLSequenceValue

ObjectIdentifierType ::= OBJECT IDENTIFIER

ObjectIdentifierValue ::= " { " ObjIdComponentsList " } " | " { " DefinedValue ObjIdComponentsList " } "

ObjIdComponentsList ::=
		ObjIdComponents
		| ObjIdComponents ObjIdComponentsList

ObjIdComponents ::= NameForm | NumberForm | NameAndNumberForm | DefinedValue

NameForm ::= identifier

NumberForm ::= number | DefinedValue

NameAndNumberForm ::= identifier " ( " NumberForm " ) "

XMLObjectIdentifierValue ::= XMLObjIdComponentList

XMLObjIdComponentList ::= XMLObjIdComponent | XMLObjIdComponent & " . " & XMLObjIdComponentList

XMLObjIdComponent ::= NameForm | XMLNumberForm | XMLNameAndNumberForm

XMLNumberForm ::= number

XMLNameAndNumberForm ::= identifier & " ( " & XMLNumberForm & " ) "

RelativeOIDType ::= RELATIVE-OID

RelativeOIDValue ::= " { " RelativeOIDComponentsList " } "

RelativeOIDComponentsList ::= RelativeOIDComponents | RelativeOIDComponents RelativeOIDComponentsList

RelativeOIDComponents ::= NumberForm | NameAndNumberForm | DefinedValue

XMLRelativeOIDValue ::= XMLRelativeOIDComponentList

XMLRelativeOIDComponentList ::= XMLRelativeOIDComponent | XMLRelativeOIDComponent & " . " & XMLRelativeOIDComponentList

XMLRelativeOIDComponent ::= XMLNumberForm | XMLNameAndNumberForm

CharacterStringType ::= RestrictedCharacterStringType | UnrestrictedCharacterStringType

RestrictedCharacterStringType ::=
		BMPString
		| GeneralString
		| GraphicString
		| IA5String
		| ISO646String
		| NumericString
		| PrintableString
		| TeletexString
		| T61String
		| UniversalString
		| UTF8String
		| VideotexString
		| VisibleString

RestrictedCharacterStringValue ::= cstring | CharacterStringList | Quadruple | Tuple

CharacterStringList ::= " { " CharSyms " } "

CharSyms ::= CharsDefn | CharSyms "," CharsDefn

CharsDefn ::= cstring | Quadruple | Tuple | DefinedValue

Quadruple ::= " { " Group "," Plane "," Row "," Cell " } "

Group ::= number
Plane ::= number
Row ::= number
Cell ::= number

Tuple ::= " { " TableColumn "," TableRow " } "

TableColumn ::= number

TableRow ::= number

XMLRestrictedCharacterStringValue ::= xmlcstring

UnrestrictedCharacterStringType ::= CHARACTER STRING
CharacterStringValue ::= RestrictedCharacterStringValue | UnrestrictedCharacterStringValue

XMLCharacterStringValue ::= XMLRestrictedCharacterStringValue |XMLUnrestrictedCharacterStringValue

UnrestrictedCharacterStringValue ::= SequenceValue

XMLUnrestrictedCharacterStringValue ::= XMLSequenceValue

UsefulType ::= typereference

/*






*/
ConstrainedType ::= Type Constraint | TypeWithConstraint

TypeWithConstraint ::=
		SET Constraint OF Type
		| SET SizeConstraint OF Type
		| SEQUENCE Constraint OF Type
		| SEQUENCE SizeConstraint OF Type
		| SET Constraint OF NamedType
		| SET SizeConstraint OF NamedType
		| SEQUENCE Constraint OF NamedType
		| SEQUENCE SizeConstraint OF NamedType

Constraint ::= " ( " ConstraintSpec ExceptionSpec " ) "

ConstraintSpec ::= SubtypeConstraint | GeneralConstraint

ExceptionSpec ::= " ! " ExceptionIdentification | empty

ExceptionIdentification ::= SignedNumber | DefinedValue | Type " : " Value

SubtypeConstraint ::= ElementSetSpecs

ElementSetSpecs ::=
		RootElementSetSpec
		| RootElementSetSpec "," " ... "
		| RootElementSetSpec "," " ... " "," AdditionalElementSetSpec

RootElementSetSpec ::= ElementSetSpec

AdditionalElementSetSpec ::= ElementSetSpec

ElementSetSpec ::= Unions | ALL Exclusions

Unions ::= Intersections | UElems UnionMark Intersections

UElems ::= Unions

Intersections ::= IntersectionElements | IElems IntersectionMark IntersectionElements

IElems ::= Intersections

IntersectionElements ::= Elements | Elems Exclusions

Elems ::= Elements

Exclusions ::= EXCEPT Elements

UnionMark ::= "|" | UNION

IntersectionMark ::= " ^ " | INTERSECTION

Elements ::= SubtypeElements | ObjectSetElements | " ( " ElementSetSpec " ) "

SubtypeElements ::=
		SingleValue
		| ContainedSubtype
		| ValueRange
		| PermittedAlphabet
		| SizeConstraint
		| TypeConstraint
		| InnerTypeConstraints
		| PatternConstraint

SingleValue ::= Value

ContainedSubtype ::= Includes Type

Includes ::= INCLUDES | empty

ValueRange ::= LowerEndpoint " .. " UpperEndpoint

LowerEndpoint ::= LowerEndValue | LowerEndValue "<"
UpperEndpoint ::= UpperEndValue | "<" UpperEndValue

LowerEndValue ::= Value | MIN

UpperEndValue ::= Value | MAX

SizeConstraint ::= SIZE Constraint

PermittedAlphabet ::= FROM Constraint

TypeConstraint ::= Type
InnerTypeConstraints ::= WITH COMPONENT SingleTypeConstraint | WITH COMPONENTS MultipleTypeConstraints

SingleTypeConstraint::= Constraint

MultipleTypeConstraints ::= FullSpecification | PartialSpecification

FullSpecification ::= " { " TypeConstraints " } "

PartialSpecification ::= " { " " ... " "," TypeConstraints " } "

TypeConstraints ::= NamedConstraint | NamedConstraint "," TypeConstraints

NamedConstraint ::= identifier ComponentConstraint

ComponentConstraint ::= ValueConstraint PresenceConstraint

ValueConstraint ::= Constraint | empty

PresenceConstraint ::= PRESENT | ABSENT | OPTIONAL | empty

PatternConstraint ::= PATTERN Value

/* grammar from itu-t X-681-2002-07 */
DefinedObjectClass ::= ExternalObjectClassReference | objectclassreference | UsefulObjectClassReference

ExternalObjectClassReference ::= modulereference " . " objectclassreference

UsefulObjectClassReference ::= TYPE-IDENTIFIER | ABSTRACT-SYNTAX

ObjectClassAssignment ::= objectclassreference "::=" ObjectClass

ObjectClass ::= DefinedObjectClass | ObjectClassDefn | ParameterizedObjectClass

ObjectClassDefn ::= CLASS " { " FieldSpec "," + " } " WithSyntaxSpec?

FieldSpec ::=
		TypeFieldSpec
		| FixedTypeValueFieldSpec
		| VariableTypeValueFieldSpec
		| FixedTypeValueSetFieldSpec
		| VariableTypeValueSetFieldSpec
		| ObjectFieldSpec
		| ObjectSetFieldSpec

PrimitiveFieldName ::=
		typefieldreference
		| valuefieldreference
		| valuesetfieldreference
		| objectfieldreference
		| objectsetfieldreference

FieldName ::= PrimitiveFieldName " . " +

TypeFieldSpec ::= typefieldreference TypeOptionalitySpec?

TypeOptionalitySpec ::= OPTIONAL | DEFAULT Type

FixedTypeValueFieldSpec ::= valuefieldreference Type UNIQUE ? ValueOptionalitySpec ?

ValueOptionalitySpec ::= OPTIONAL | DEFAULT Value

VariableTypeValueFieldSpec ::= valuefieldreference FieldName ValueOptionalitySpec ?

FixedTypeValueSetFieldSpec ::= valuesetfieldreference Type ValueSetOptionalitySpec ?

ValueSetOptionalitySpec ::= OPTIONAL | DEFAULT ValueSet
VariableTypeValueSetFieldSpec ::= valuesetfieldreference FieldName ValueSetOptionalitySpec?

ObjectFieldSpec ::= objectfieldreference DefinedObjectClass ObjectOptionalitySpec?

ObjectOptionalitySpec ::= OPTIONAL | DEFAULT Object

ObjectSetFieldSpec ::= objectsetfieldreference DefinedObjectClass ObjectSetOptionalitySpec ?

ObjectSetOptionalitySpec ::= OPTIONAL | DEFAULT ObjectSet

WithSyntaxSpec ::= WITH SYNTAX SyntaxList

SyntaxList ::= " { " TokenOrGroupSpec empty + " } "

TokenOrGroupSpec ::= RequiredToken | OptionalGroup

OptionalGroup ::= "[" TokenOrGroupSpec empty + "]"

RequiredToken ::= Literal | PrimitiveFieldName

Literal ::= word | ","

DefinedObject ::= ExternalObjectReference | objectreference

ExternalObjectReference ::= modulereference " . " objectreference

ObjectAssignment ::= objectreference DefinedObjectClass " ::= " Object

Object ::=
		DefinedObject
		| ObjectDefn
		| ObjectFromObject
		| ParameterizedObject

ObjectDefn ::= DefaultSyntax | DefinedSyntax

DefaultSyntax ::= " { " FieldSetting "," * " } "

FieldSetting ::= PrimitiveFieldName Setting

DefinedSyntax ::= " { " DefinedSyntaxToken empty * " } "

DefinedSyntaxToken ::= Literal | Setting

Setting ::= Type | Value | ValueSet | Object | ObjectSet

DefinedObjectSet ::= ExternalObjectSetReference | objectsetreference

ExternalObjectSetReference ::= modulereference " . " objectsetreference

ObjectSetAssignment ::= objectsetreference DefinedObjectClass "::=" ObjectSet

ObjectSet ::= " { " ObjectSetSpec " } "

ObjectSetSpec ::=
		RootElementSetSpec
		| RootElementSetSpec "," " ... "
		| " ... "
		| " ... " "," AdditionalElementSetSpec
		| RootElementSetSpec "," " ... " "," AdditionalElementSetSpec

ObjectSetElements ::= Object | DefinedObjectSet | ObjectSetFromObjects | ParameterizedObjectSet

ObjectClassFieldType ::= DefinedObjectClass " . " FieldName

ObjectClassFieldValue ::= OpenTypeFieldVal | FixedTypeFieldVal

OpenTypeFieldVal ::= Type " : " Value

FixedTypeFieldVal ::= BuiltinValue | ReferencedValue

XMLObjectClassFieldValue ::= XMLOpenTypeFieldVal | XMLFixedTypeFieldVal

XMLOpenTypeFieldVal ::= XMLTypedValue

XMLFixedTypeFieldVal ::= XMLBuiltinValue

InformationFromObjects ::=
		ValueFromObject
		| ValueSetFromObjects
		| TypeFromObject
		| ObjectFromObject
		| ObjectSetFromObjects

ReferencedObjects ::= DefinedObject | ParameterizedObject | DefinedObjectSet | ParameterizedObjectSet

ValueFromObject ::= ReferencedObjects " . " FieldName

ValueSetFromObjects ::= ReferencedObjects " . " FieldName

TypeFromObject ::= ReferencedObjects " . " FieldName

ObjectFromObject ::= ReferencedObjects " . " FieldName

ObjectSetFromObjects ::= ReferencedObjects " . " FieldName

InstanceOfType ::= INSTANCE OF DefinedObjectClass

InstanceOfValue ::= Value

XMLInstanceOfValue ::= XMLValue

%%
