package helpers

import (
	"testing"
)

func TestBezier(t *testing.T) {

	type value struct {
		start, ctrl1, ctrl2, end [2]float64
		lowBound, heighBound     float64
	}

	cases := []struct {
		value
		expected map[int]float64
	}{
		{
			value{
				[2]float64{20, 0},
				[2]float64{90, 120},
				[2]float64{186, 144},
				[2]float64{255, 230},
				0,
				255,
			},
			map[int]float64{0: 0.047619047619047616, 1: 0.09523809523809523, 2: 0.14285714285714285, 3: 0.19047619047619047, 4: 0.23809523809523808, 5: 0.2857142857142857, 6: 0.3333333333333333, 7: 0.38095238095238093, 8: 0.42857142857142855, 9: 0.47619047619047616, 10: 0.5238095238095238, 11: 0.5714285714285714, 12: 0.6190476190476191, 13: 0.6666666666666667, 14: 0.7142857142857143, 15: 0.7619047619047619, 16: 0.8095238095238095, 17: 0.8571428571428572, 18: 0.9047619047619048, 19: 0.9523809523809523, 20: 1, 21: 3, 22: 4, 23: 6, 24: 7, 25: 9, 26: 11, 27: 12, 28: 14, 29: 15, 30: 17, 31: 18, 32: 20, 33: 21, 34: 23, 35: 24, 36: 26, 37: 27, 38: 29, 39: 30, 40: 31, 41: 33, 42: 34, 43: 35, 44: 37, 45: 38, 46: 40, 47: 41, 48: 42, 49: 43, 50: 45, 51: 46, 52: 47, 53: 48, 54: 50, 55: 51, 56: 52, 57: 53, 58: 54, 59: 56, 60: 57, 61: 58, 62: 59, 63: 60, 64: 61, 65: 63, 66: 64, 67: 65, 68: 66, 69: 67, 70: 68, 71: 69, 72: 70, 73: 71, 74: 72, 75: 73, 76: 74, 77: 75, 78: 76, 79: 77, 80: 78, 81: 79, 82: 80, 83: 81, 84: 82, 85: 83, 86: 84, 87: 85, 88: 86, 89: 87, 90: 88, 91: 89, 92: 90, 93: 91, 94: 92, 95: 93, 96: 94, 97: 94, 98: 95, 99: 96, 100: 97, 101: 98, 102: 99, 103: 100, 104: 101, 105: 102, 106: 102, 107: 103, 108: 104, 109: 105, 110: 106, 111: 107, 112: 107, 113: 108, 114: 109, 115: 110, 116: 111, 117: 111, 118: 112, 119: 113, 120: 114, 121: 115, 122: 116, 123: 116, 124: 117, 125: 118, 126: 119, 127: 120, 128: 120, 129: 121, 130: 122, 131: 123, 132: 124, 133: 124, 134: 125, 135: 126, 136: 127, 137: 127, 138: 128, 139: 129, 140: 130, 141: 130, 142: 131, 143: 132, 144: 133, 145: 133, 146: 134, 147: 135, 148: 136, 149: 136, 150: 137, 151: 138, 152: 139, 153: 140, 154: 140, 155: 141, 156: 142, 157: 143, 158: 143, 159: 144, 160: 145, 161: 146, 162: 146, 163: 147, 164: 148, 165: 148, 166: 149, 167: 150, 168: 151, 169: 151, 170: 152, 171: 153, 172: 154, 173: 155, 174: 155, 175: 156, 176: 157, 177: 158, 178: 158, 179: 159, 180: 160, 181: 161, 182: 161, 183: 162, 184: 163, 185: 164, 186: 165, 187: 165, 188: 166, 189: 167, 190: 168, 191: 168, 192: 169, 193: 170, 194: 171, 195: 172, 196: 173, 197: 173, 198: 174, 199: 175, 200: 176, 201: 177, 202: 177, 203: 178, 204: 179, 205: 180, 206: 181, 207: 182, 208: 183, 209: 183, 210: 184, 211: 185, 212: 186, 213: 187, 214: 188, 215: 189, 216: 189, 217: 191, 218: 191, 219: 192, 220: 193, 221: 194, 222: 195, 223: 196, 224: 197, 225: 198, 226: 199, 227: 200, 228: 201, 229: 202, 230: 203, 231: 204, 232: 205, 233: 206, 234: 207, 235: 208, 236: 209, 237: 210, 238: 211, 239: 212, 240: 213, 241: 214, 242: 215, 243: 216, 244: 218, 245: 219, 246: 220, 247: 221, 248: 222, 249: 223, 250: 224, 251: 226, 252: 227, 253: 228, 254: 229, 255: 230},
		},
		{
			value{
				[2]float64{0, 0},
				[2]float64{144, 90},
				[2]float64{138, 120},
				[2]float64{255, 255},
				0,
				255,
			},
			map[int]float64{0: 0, 1: 1, 2: 1, 3: 2, 4: 3, 5: 3, 6: 4, 7: 5, 8: 5, 9: 6, 10: 6, 11: 7, 12: 8, 13: 8, 14: 9, 15: 10, 16: 10, 17: 11, 18: 12, 19: 12, 20: 13, 21: 14, 22: 14, 23: 15, 24: 16, 25: 16, 26: 17, 27: 18, 28: 18, 29: 19, 30: 20, 31: 20, 32: 21, 33: 22, 34: 22, 35: 23, 36: 23, 37: 24, 38: 25, 39: 26, 40: 26, 41: 27, 42: 28, 43: 28, 44: 29, 45: 30, 46: 30, 47: 31, 48: 32, 49: 33, 50: 33, 51: 34, 52: 35, 53: 35, 54: 36, 55: 37, 56: 37, 57: 38, 58: 39, 59: 40, 60: 40, 61: 41, 62: 42, 63: 43, 64: 44, 65: 44, 66: 45, 67: 46, 68: 46, 69: 47, 70: 48, 71: 49, 72: 49, 73: 50, 74: 51, 75: 52, 76: 53, 77: 53, 78: 54, 79: 55, 80: 56, 81: 57, 82: 57, 83: 58, 84: 59, 85: 60, 86: 61, 87: 62, 88: 62, 89: 63, 90: 64, 91: 65, 92: 66, 93: 67, 94: 67, 95: 68, 96: 69, 97: 70, 98: 71, 99: 72, 100: 73, 101: 74, 102: 75, 103: 75, 104: 76, 105: 77, 106: 78, 107: 79, 108: 80, 109: 81, 110: 82, 111: 83, 112: 84, 113: 85, 114: 86, 115: 87, 116: 88, 117: 89, 118: 90, 119: 91, 120: 92, 121: 93, 122: 94, 123: 95, 124: 96, 125: 97, 126: 98, 127: 99, 128: 100, 129: 102, 130: 103, 131: 104, 132: 105, 133: 106, 134: 107, 135: 108, 136: 109, 137: 110, 138: 111, 139: 113, 140: 114, 141: 115, 142: 116, 143: 117, 144: 119, 145: 120, 146: 121, 147: 122, 148: 123, 149: 125, 150: 126, 151: 127, 152: 128, 153: 129, 154: 131, 155: 132, 156: 133, 157: 134, 158: 136, 159: 137, 160: 138, 161: 139, 162: 141, 163: 142, 164: 143, 165: 145, 166: 146, 167: 147, 168: 148, 169: 149, 170: 151, 171: 152, 172: 153, 173: 155, 174: 156, 175: 157, 176: 158, 177: 160, 178: 161, 179: 162, 180: 164, 181: 165, 182: 166, 183: 168, 184: 169, 185: 170, 186: 171, 187: 173, 188: 174, 189: 175, 190: 177, 191: 178, 192: 179, 193: 180, 194: 181, 195: 183, 196: 184, 197: 185, 198: 186, 199: 188, 200: 189, 201: 190, 202: 192, 203: 193, 204: 194, 205: 195, 206: 197, 207: 198, 208: 199, 209: 200, 210: 201, 211: 203, 212: 204, 213: 205, 214: 207, 215: 208, 216: 209, 217: 210, 218: 211, 219: 213, 220: 214, 221: 215, 222: 216, 223: 217, 224: 219, 225: 220, 226: 221, 227: 222, 228: 223, 229: 225, 230: 226, 231: 227, 232: 229, 233: 230, 234: 231, 235: 232, 236: 233, 237: 234, 238: 236, 239: 237, 240: 238, 241: 239, 242: 240, 243: 242, 244: 243, 245: 244, 246: 245, 247: 246, 248: 247, 249: 249, 250: 250, 251: 251, 252: 252, 253: 253, 254: 254, 255: 255},
		},
		{
			value{
				[2]float64{10, 0},
				[2]float64{115, 105},
				[2]float64{148, 100},
				[2]float64{255, 248},
				0,
				255,
			},
			map[int]float64{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0, 10: 0, 11: 1, 12: 2, 13: 3, 14: 4, 15: 5, 16: 6, 17: 7, 18: 8, 19: 9, 20: 10, 21: 11, 22: 12, 23: 13, 24: 14, 25: 15, 26: 16, 27: 17, 28: 18, 29: 19, 30: 20, 31: 21, 32: 22, 33: 23, 34: 24, 35: 25, 36: 26, 37: 26, 38: 27, 39: 28, 40: 29, 41: 30, 42: 31, 43: 32, 44: 33, 45: 34, 46: 35, 47: 36, 48: 37, 49: 38, 50: 38, 51: 39, 52: 40, 53: 41, 54: 42, 55: 43, 56: 44, 57: 45, 58: 46, 59: 46, 60: 47, 61: 48, 62: 49, 63: 50, 64: 51, 65: 52, 66: 53, 67: 54, 68: 54, 69: 55, 70: 56, 71: 57, 72: 58, 73: 59, 74: 60, 75: 60, 76: 61, 77: 62, 78: 63, 79: 64, 80: 65, 81: 66, 82: 66, 83: 67, 84: 68, 85: 69, 86: 70, 87: 71, 88: 71, 89: 72, 90: 73, 91: 74, 92: 75, 93: 76, 94: 76, 95: 77, 96: 78, 97: 79, 98: 80, 99: 81, 100: 81, 101: 82, 102: 83, 103: 84, 104: 85, 105: 86, 106: 86, 107: 87, 108: 88, 109: 89, 110: 90, 111: 91, 112: 91, 113: 92, 114: 93, 115: 94, 116: 95, 117: 96, 118: 96, 119: 97, 120: 98, 121: 99, 122: 100, 123: 101, 124: 102, 125: 102, 126: 103, 127: 104, 128: 105, 129: 106, 130: 107, 131: 108, 132: 108, 133: 109, 134: 110, 135: 111, 136: 112, 137: 113, 138: 114, 139: 115, 140: 115, 141: 116, 142: 117, 143: 118, 144: 119, 145: 120, 146: 121, 147: 122, 148: 123, 149: 124, 150: 125, 151: 126, 152: 127, 153: 128, 154: 129, 155: 129, 156: 131, 157: 132, 158: 132, 159: 133, 160: 134, 161: 135, 162: 136, 163: 137, 164: 138, 165: 139, 166: 140, 167: 141, 168: 142, 169: 143, 170: 145, 171: 145, 172: 147, 173: 148, 174: 149, 175: 150, 176: 151, 177: 152, 178: 153, 179: 154, 180: 155, 181: 156, 182: 157, 183: 158, 184: 159, 185: 161, 186: 162, 187: 163, 188: 164, 189: 165, 190: 166, 191: 167, 192: 168, 193: 169, 194: 171, 195: 172, 196: 173, 197: 174, 198: 175, 199: 176, 200: 178, 201: 179, 202: 180, 203: 181, 204: 182, 205: 184, 206: 185, 207: 186, 208: 187, 209: 188, 210: 190, 211: 191, 212: 192, 213: 193, 214: 194, 215: 196, 216: 197, 217: 198, 218: 200, 219: 201, 220: 202, 221: 203, 222: 205, 223: 206, 224: 207, 225: 208, 226: 210, 227: 211, 228: 212, 229: 213, 230: 215, 231: 216, 232: 218, 233: 219, 234: 220, 235: 221, 236: 223, 237: 224, 238: 225, 239: 227, 240: 228, 241: 229, 242: 231, 243: 232, 244: 233, 245: 235, 246: 236, 247: 238, 248: 239, 249: 240, 250: 241, 251: 243, 252: 244, 253: 246, 254: 247, 255: 248},
		},
		{
			value{
				[2]float64{0, 0},
				[2]float64{120, 100},
				[2]float64{120, 140},
				[2]float64{255, 255},
				0,
				255,
			},
			map[int]float64{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 4, 6: 5, 7: 6, 8: 7, 9: 8, 10: 9, 11: 10, 12: 10, 13: 11, 14: 12, 15: 13, 16: 14, 17: 15, 18: 16, 19: 17, 20: 17, 21: 18, 22: 19, 23: 20, 24: 21, 25: 22, 26: 23, 27: 23, 28: 25, 29: 25, 30: 26, 31: 27, 32: 28, 33: 29, 34: 30, 35: 31, 36: 32, 37: 33, 38: 34, 39: 34, 40: 35, 41: 36, 42: 37, 43: 38, 44: 39, 45: 40, 46: 41, 47: 42, 48: 43, 49: 44, 50: 45, 51: 46, 52: 47, 53: 48, 54: 49, 55: 50, 56: 51, 57: 52, 58: 53, 59: 54, 60: 55, 61: 56, 62: 57, 63: 58, 64: 59, 65: 60, 66: 61, 67: 62, 68: 63, 69: 64, 70: 65, 71: 66, 72: 67, 73: 68, 74: 69, 75: 70, 76: 71, 77: 72, 78: 73, 79: 74, 80: 75, 81: 76, 82: 77, 83: 78, 84: 79, 85: 81, 86: 82, 87: 83, 88: 84, 89: 85, 90: 86, 91: 87, 92: 88, 93: 89, 94: 91, 95: 92, 96: 93, 97: 94, 98: 95, 99: 96, 100: 97, 101: 98, 102: 99, 103: 101, 104: 102, 105: 103, 106: 104, 107: 105, 108: 106, 109: 108, 110: 109, 111: 110, 112: 111, 113: 112, 114: 113, 115: 114, 116: 115, 117: 117, 118: 118, 119: 119, 120: 120, 121: 121, 122: 123, 123: 124, 124: 125, 125: 126, 126: 127, 127: 128, 128: 129, 129: 131, 130: 132, 131: 133, 132: 134, 133: 135, 134: 136, 135: 138, 136: 139, 137: 140, 138: 141, 139: 142, 140: 143, 141: 144, 142: 146, 143: 146, 144: 148, 145: 149, 146: 150, 147: 151, 148: 152, 149: 153, 150: 154, 151: 156, 152: 156, 153: 158, 154: 159, 155: 160, 156: 161, 157: 162, 158: 163, 159: 164, 160: 165, 161: 166, 162: 167, 163: 168, 164: 170, 165: 171, 166: 172, 167: 173, 168: 174, 169: 175, 170: 176, 171: 177, 172: 178, 173: 179, 174: 180, 175: 181, 176: 182, 177: 183, 178: 184, 179: 185, 180: 186, 181: 187, 182: 188, 183: 189, 184: 190, 185: 191, 186: 192, 187: 193, 188: 194, 189: 195, 190: 196, 191: 197, 192: 198, 193: 199, 194: 200, 195: 201, 196: 202, 197: 203, 198: 204, 199: 205, 200: 206, 201: 207, 202: 208, 203: 209, 204: 210, 205: 210, 206: 212, 207: 212, 208: 213, 209: 214, 210: 215, 211: 216, 212: 217, 213: 218, 214: 219, 215: 220, 216: 221, 217: 222, 218: 223, 219: 224, 220: 224, 221: 225, 222: 226, 223: 227, 224: 228, 225: 229, 226: 230, 227: 231, 228: 232, 229: 233, 230: 233, 231: 234, 232: 235, 233: 236, 234: 237, 235: 238, 236: 239, 237: 240, 238: 241, 239: 242, 240: 242, 241: 243, 242: 244, 243: 245, 244: 246, 245: 247, 246: 248, 247: 249, 248: 249, 249: 250, 250: 251, 251: 252, 252: 253, 253: 254, 254: 254, 255: 255},
		},
	}

	for i, c := range cases {
		result := Bezier(c.start, c.ctrl1, c.ctrl2, c.end, c.lowBound, c.heighBound)
		if !MapsEqual(result, c.expected) {
			t.Errorf("%s:case number %d\nexpected:\n\t%v\nactual:\n\t%v", "Bezier", i+1, c.expected, result)
		}
	}
}

func TestMissingValues(t *testing.T) {
	type value struct {
		val  map[int]float64
		endX int
	}

	cases := []*struct {
		*value
		expected map[int]float64
	}{
		{
			&value{
				map[int]float64{1: 1, 2: 2, 3: 3},
				2,
			},
			map[int]float64{1: 1, 2: 2, 3: 3},
		},
		{
			&value{
				map[int]float64{1: 1, 2: 2, 3: 3},
				3,
			},
			map[int]float64{0: 0.5, 1: 1, 2: 2, 3: 3},
		},
		{
			&value{
				map[int]float64{1: 1, 2: 2, 3: 3, 4: 4},
				3,
			},
			map[int]float64{1: 1, 2: 2, 3: 3, 4: 4},
		},
		{
			&value{
				map[int]float64{1: 1, 2: 2, 3: 3},
				4,
			},
			map[int]float64{},
		},
	}

	for i, c := range cases {
		result := MissingValues(c.val, c.endX)
		if !MapsEqual(result, c.expected) {
			t.Errorf("%s: case number %d\nexpected: %v\nactual: %v", "MissingValues", i+1, c.expected, result)
		}
	}
}

func TestMapsEqual(t *testing.T) {
	cases := []struct {
		map1, map2 map[int]float64
		expected   bool
	}{
		{
			map[int]float64{1: 1, 2: 2, 3: 3},
			map[int]float64{1: 1, 2: 2, 3: 3},
			true,
		},
		{
			map[int]float64{1: 0, 3: 5, 2: 4},
			map[int]float64{1: 0, 2: 4, 3: 5},
			true,
		},
		{
			map[int]float64{1: 1, 2: 2},
			map[int]float64{1: 1, 2: 2, 3: 3},
			false,
		},
	}

	for i, c := range cases {
		if MapsEqual(c.map1, c.map2) != c.expected {
			t.Errorf("%s: case number: %d\nExpected:\n\t%v\nActual:\n\t%v", "MapsEqual", i+1, c.expected, MapsEqual(c.map1, c.map2))
		}
	}
}
