import java.io.File;
import java.io.FileNotFoundException;
import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;
import java.util.SortedSet;
import java.util.TreeSet;
import java.util.Iterator;
import java.util.Arrays;

public class Prob7Part1 {
    public static void main(String args[]) throws FileNotFoundException {
        SortedSet<Hand> h = new TreeSet<Hand>();
        File f = new File("input.txt");
        Scanner sc = new Scanner(f);
        while(sc.hasNextLine()) {
            String[] line = sc.nextLine().split(" ");
            var x = new Hand(line[0].toCharArray(), Integer.parseInt(line[1]));
            h.add(x);
        }
        long total = 0;
        int i = 1;
        for(Hand x: h) {
            System.out.println(x);
            total += i * x.bid;
            i++;
        }
        System.out.println(total);
    }
}

class Hand implements Comparable<Hand> {
    char[] hand;
    int bid;
    int type;
    // 6 = 5oaK, 0 = high card

    public Hand(char[] hand, int bid) {
        this.hand = hand;
        this.bid = bid;
        type = getType(hand);
    }

    public int getType(char[] hand) {
        Map<Character, Integer> cards = new HashMap<Character, Integer>();
        for(char c: hand) {
            if(!cards.containsKey(c)) {
                cards.put(c, 1);
            } else {
                cards.put(c, cards.get(c) + 1);
            }
        }

        if(cards.containsValue(5)) {
            return 6;
        } else if (cards.containsValue(4)) {
            return 5;
        } else if(cards.containsValue(3) && cards.containsValue(2)) {
            return 4;
        } else if(cards.containsValue(3)) {
            return 3;
        } else if(cards.containsValue(2) && cards.keySet().size() == 3) {
            return 2;
        } else if(cards.containsValue(2)) {
            return 1;
        }
        return 0;
    }

    public int compareTo(Hand h) {
        if(this.type != h.type) {
            return this.type- h.type;
        }
        for(int i = 0; i < hand.length; i++) {
            if(hand[i] != h.hand[i]) {
                return mapChar(hand[i]) - mapChar(h.hand[i]);
            }
        }
        return 0;
    }
    
    public boolean equals(Object o) {
        Hand h = (Hand) o;
        return compareTo(h) == 0;
    }
    
    public String toString() {
        return Arrays.toString(hand) + " " + bid + " " + type;
    }
    
    public int mapChar(char c) {
        if(c == 'A') {
            return 99;
        }
        if(c == 'K') {
            return 98;
        }
        if(c == 'Q') {
            return 97;
        }
        if(c == 'J') {
            return 96;
        }
        if(c == 'T') {
            return 95;
        }
        return c - '9';
    }
}
