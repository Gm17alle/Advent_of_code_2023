import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;
import java.io.File;
import java.io.FileNotFoundException;
import java.math.BigInteger;
import java.util.ArrayList;
import java.util.List;

/**
 * Write a description of class Part1 here.
 *
 * @author (your name)
 * @version (a version number or a date)
 */
public class Part2
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
        var curKeyList = new ArrayList<String>();
        for(var ks: map.keySet()) {
            if(ks.charAt(2) == 'A') {
                curKeyList.add(ks);
            }
        }
        var answers = new ArrayList<Integer>();
        var count = new BigInteger("0");
        var one = new BigInteger("1");
        var oneMillion = new BigInteger("1000000");
        var isDone = false;
        while(curKeyList.size() > 0) {
            int index = instructions.charAt(count.mod(new BigInteger("" + instructions.length())).intValue()) == 'L' ? 0 : 1;
            var newKeyList = new ArrayList<String>(curKeyList.size());
            isDone = true;
            count = count.add(one);
            for(var s: curKeyList) {
                String toAdd = map.get(s).getFromInt(index);
                if(toAdd.charAt(2) == 'Z') {
                    answers.add(count.intValue());
                } else {
                    newKeyList.add(toAdd);                 
                }
            }

            if(count.mod(oneMillion).intValue() == 0) {
                System.out.println("Count: " + count + " , KeyList: " + curKeyList);
            }
            curKeyList = newKeyList;

        }
        System.out.println("Answers is = " + answers + "\nThe count is " + findLCM(answers));
    }

    public static int gcd(int a, int b) {
        if (b == 0) {
            return a;
        } else {
            return gcd(b, a % b);
        }
    }

    // Function to find the LCM of two numbers
    public static int lcm(int a, int b) {
        return (a * b) / gcd(a, b);
    }

    // Function to find the LCM of a list of integers
    public static int findLCM(List<Integer> numbers) {
        if (numbers == null || numbers.isEmpty()) {
            throw new IllegalArgumentException("List cannot be null or empty");
        }

        int lcmResult = numbers.get(0);
        for (int i = 1; i < numbers.size(); i++) {
            int currentNumber = numbers.get(i);
            lcmResult = lcm(lcmResult, currentNumber);
        }
        return lcmResult;
    }
}

class Pair {
    String left;
    String right;
    public Pair(String l, String r) {
        this.left = l;
        this.right = r;
    }

    public String getFromInt(int i) {
        if(i == 0) {
            return left;
        }
        return right;
    }
}
