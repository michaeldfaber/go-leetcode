func suggestedProducts(products []string, searchWord string) [][]string {
    var searchResults [][]string
    sort.Strings(products)
    
    for i := 1; i <= len(searchWord); i++ {
        var searchResult []string
        for j := 0; j < len(products); j++ {
            if len(products[j]) >= i {
                if searchWord[0:i] == products[j][0:i] {
                    searchResult = append(searchResult, products[j])
                    if len(searchResult) == 3 {
                        break
                    } else {
                        continue
                    }
                }
            }
            products = append(products[:j], products[j+1:]...)
            if j == 0 {
                j--
            }
        }
        searchResults = append(searchResults, searchResult)
    }
    
    return searchResults
}