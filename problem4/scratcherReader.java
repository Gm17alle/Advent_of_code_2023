import java.io.File;
import java.util.Scanner;
import java.util.HashSet;
import java.util.Set;

/**
 * Write a description of class scratcherReader here.
 *
 * @author (your name)
 * @version (a version number or a date)
 */
public class scratcherReader
{
    public static void main(String args[]) {
        Scanner sc = null;
        try {
            File f = new File("cards.txt");
            sc = new Scanner(f);
        } catch (Throwable t) {
            // do nothing
        }

        int[] countWins = new int[203];
        for(int i = 0; i < 203; i++) {
            countWins[i] = 1;
        }
        int sum = 0;
        int points = 0;
        int index = 0;
        while(sc.hasNext()) {
            int curWins = 0;
            String line = sc.nextLine();
            String everything = line.split(":")[1];
            Set<Integer> winningNums = new HashSet<Integer>();
            var p = everything.split("\\|")[0].trim();
            var q = everything.split("\\|")[1].trim();
            Scanner getNums = new Scanner(p);
            while(getNums.hasNextInt()) {
                winningNums.add(getNums.nextInt());
            }
            Scanner theNums = new Scanner(q);
            while(theNums.hasNextInt()) {
                if(winningNums.contains(theNums.nextInt())) {
                    curWins++;
                }
            }

            for(int i = 0; i < countWins[index]; i++) {
                for(int j = 1; j <= curWins; j++) {
                    if(index+j<countWins.length) {
                        countWins[index+j] = countWins[index+j] + 1;    
                    }
                }
            }
            points += Math.pow(2,curWins-1);
            index++;

        }
        for(int numCards: countWins) {
            sum += numCards;
        }
        System.out.printf("The amount of part 1 points is %d and the number of cards is %d", points, sum);
    }
}
