package joker

// CardFace defines a card face
type CardFace int

const (
	FaceAce   CardFace = 1
	Face2     CardFace = 2
	Face3     CardFace = 3
	Face4     CardFace = 4
	Face5     CardFace = 5
	Face6     CardFace = 6
	Face7     CardFace = 7
	Face8     CardFace = 8
	Face9     CardFace = 9
	Face10    CardFace = 10
	FaceJack  CardFace = 11
	FaceQueen CardFace = 12
	FaceKing  CardFace = 13
	FaceJoker CardFace = 14
)

// StandardFaces is a slice of all faces except Jokers.
var StandardFaces = []CardFace{
	FaceAce,
	Face2,
	Face3,
	Face4,
	Face5,
	Face6,
	Face7,
	Face8,
	Face9,
	Face10,
	FaceJack,
	FaceQueen,
	FaceKing,
}

// AllFaces is a slice of all faces.
var AllFaces = []CardFace{
	FaceAce,
	Face2,
	Face3,
	Face4,
	Face5,
	Face6,
	Face7,
	Face8,
	Face9,
	Face10,
	FaceJack,
	FaceQueen,
	FaceKing,
	FaceJoker,
}

// Name returns the card face name.
func (f CardFace) Name() string {
	switch f {
	case FaceAce:
		return "Ace"
	case Face2:
		return "2"
	case Face3:
		return "3"
	case Face4:
		return "4"
	case Face5:
		return "5"
	case Face6:
		return "6"
	case Face7:
		return "7"
	case Face8:
		return "8"
	case Face9:
		return "9"
	case Face10:
		return "10"
	case FaceJack:
		return "Jack"
	case FaceQueen:
		return "Queen"
	case FaceKing:
		return "King"
	case FaceJoker:
		return "Joker"
	default:
		return "Invalid"
	}
}

// String returns the card face name.
func (f CardFace) String() string {
	return f.Name()
}
