package easy

import "strings"

func numUniqueEmails(emails []string) int {
	set := make(map[string]struct{})
	for _, email := range emails {
		realEmail := getRealMail(email)
		set[realEmail] = struct{}{}
	}

	return len(set)
}

func getRealMail(mail string) string {
	var builder strings.Builder
	inDomain := false
	for i := 0; i < len(mail); i++ {
		if mail[i] == '@' {
			inDomain = true
			builder.WriteString(mail[i:])
			break
		}

		if !inDomain {
			if mail[i] == '.' {
				continue
			}
			if mail[i] == '+' {
				for i < len(mail) && mail[i] != '@' {
					i++
				}
				i--
				continue
			}
		}

		builder.WriteByte(mail[i])
	}

	return builder.String()
}
