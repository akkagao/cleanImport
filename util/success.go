package util

import "fmt"

func ShowSuccess() {
	fmt.Println("                              ###                                                           ")
	fmt.Println("                     #      #####                                                           ")
	fmt.Println("                      #########                                                             ")
	fmt.Println("                        ###L##                                                              ")
	fmt.Println("                           ##                                                               ")
	fmt.Println("                           #: #                                                             ")
	fmt.Println("                          ## K#                                                             ")
	fmt.Println("                         ## ####                            #      ###                      ")
	fmt.Println("                         #   ##                   WW         #      #E                      ")
	fmt.Println("                        #Wt###f:                  D#          #                             ")
	fmt.Println("                       ## # #            #        ##           ##                  #        ")
	fmt.Println("                       #   #   #;f       D####     #            ##                          ")
	fmt.Println("                      ##############.     #       i#             ##  #   #        #         ")
	fmt.Println("                     ################    #  #  # #f#      ##: #.  .# #    ##K:#      i      ")
	fmt.Println("                     ####      :   ###   #L#   #  i#    ##    GD   ###      #   ;##L###     ")
	fmt.Println("                    K#    #   :W   ###     .   ## j#  ###     #    ###               :#     ")
	fmt.Println("                    ;    ##  ##    ##f     : # #t .# G#      #     # ##    #    #    ##     ")
	fmt.Println("                     t  ##  ##     ##     ##j      #        ###       #    ##   f   ##      ")
	fmt.Println("           ##   #   # ,######     K######## #      #      ##  ##    ####   ##  D   #        ")
	fmt.Println("          ##   ##  ########       ##      ###     ,#    .##  #K      W##;      #            ")
	fmt.Println("         ##  L# #### ###         ##       t#      ##    ##            :##      ##           ")
	fmt.Println("         #####         #       ###               ##                     #                   ")
	fmt.Println("          ###          ###EL#####              ##                       ,                   ")
	fmt.Println("                        #######                                       .  .                  ")
	fmt.Println("                         t##                                                                ")
}


func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}