/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    i, j := 0, n
    
    for i < j {
        k := (j + i) / 2
        isBad := isBadVersion(k)
        
        if isBad {
            j = k
        } else {
            i = k + 1
        }
    }
    
    return i
}