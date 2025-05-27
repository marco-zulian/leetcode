func numUniqueEmails(emails []string) int {
    sentEmails := map[string]bool{}
    var sent int
    
    for _, email := range emails {
        splittedEmail := strings.SplitN(email, "@", 2)
        
        local, domain := splittedEmail[0], splittedEmail[1]
        formatedLocal := interpretLocal(local)
        
        finalEmail := formatedLocal + "@" + domain
        
        if isValidDomain(domain) && !sentEmails[finalEmail] {
            sentEmails[finalEmail] = true
            sent++
        }
    }
    
    return sent
}

func interpretLocal(local string) string {
    parts := strings.SplitN(local, "+", 2)
    
    return strings.ReplaceAll(parts[0], ".", "")
}

func isValidDomain(domain string) bool {
    return strings.HasSuffix(domain, ".com") && len(domain) > 4
}