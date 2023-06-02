package service

func FindLongestSubstring(s string) string {
	// Инициализируем переменные для хранения информации о текущей подстроке
	var start, end int //:= 0, 0
	var longest string // := ""

	// Инициализируем словарь для отслеживания повторяющихся символов
	seen := make(map[byte]int)

	// Обходим строку и на каждой позиции проверяем, повторяется ли текущий символ
	for end < len(s) {
		char := s[end]
		if pos, ok := seen[char]; ok && pos >= start {
			// Если символ уже встречался в текущей подстроке, обновляем начальную позицию
			start = pos + 1
		}

		// Обновляем конечную позицию и проверяем, является ли текущая подстрока самой длинной
		seen[char] = end
		end++
		if end-start > len(longest) {
			longest = s[start:end]
		}
	}

	return longest
}
