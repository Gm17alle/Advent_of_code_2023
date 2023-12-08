import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;
import java.io.File;
import java.io.FileNotFoundException;
import java.math.BigInteger;

/**
 * Write a description of class Part1 here.
 *
 * @author (your name)
 * @version (a version number or a date)
 */
public class Part1
{
    public static void main(String args[]) throws FileNotFoundException {
        Map<String, Pair> map = new HashMap<String, Pair>();
        File f = new File("input.txt");
        Scanner sc = new Scanner(f);
        String instructions = sc.nextLine();
        sc.nextLine();
        while(sc.hasNextLine()) {
            var line = sc.nextLine().split(" = ");
            var key = line[0];
            var left = line[1].split(",")[0].substring(1,4);
            var right = line[1].split(",")[1].substring(1,4);
            map.put(key, new Pair(left, right));

        }
        var curKey = "AAA";
        var count = new BigInteger("0");
        var one = new BigInteger("1");
        var oneMillion = new BigInteger("1000000");
        while(!curKey.equals("ZZZ")) {
            if(instructions.charAt(count.mod(new BigInteger("" + instructions.length())).intValue()) == 'L') {
                curKey = map.get(curKey).left;
            } else {
                curKey = map.get(curKey).right;
            }
            if(count.mod(oneMillion).intValue() == 0) {
                System.out.println("Count: " + count + " , Key: " + curKey);
            }
            count = count.add(one);

        }
        System.out.println("The count is " + count);
    }
}

class Pair {
    String left;
    String right;
    public Pair(String l, String r) {
        this.left = l;
        this.right = r;
    }
}
