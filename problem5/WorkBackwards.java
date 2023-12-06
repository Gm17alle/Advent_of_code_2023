import java.io.File;
import java.util.Scanner;
import java.io.FileNotFoundException;
import java.util.*;

/**
 * Write a description of class WorkBackwards here.
 *
 * @author (your name)
 * @version (a version number or a date)
 */
public class WorkBackwards
{
    public static void main(String args[]) throws FileNotFoundException {
        long start = System.nanoTime();
        File f = new File("maps.txt");
        Scanner sc = new Scanner(f);
        String s = "";
        while(sc.hasNextLine()) {
            s += sc.nextLine() + "\n";
        }

        var segments = s.split(":");
        Set<Seed> seeds = new HashSet<Seed>();
        Scanner scanSeeds = new Scanner(segments[1]);
        while(scanSeeds.hasNextDouble()) {
            seeds.add(new Seed(scanSeeds.nextDouble(), scanSeeds.nextDouble()));
        }

        var rangeMaps = new ArrayList<List<R>>(); //poop var name, sorry not sorry
        for(int i = 2; i < segments.length; i++) {
            List<R> ranges = new ArrayList<R>();
            var currSc = new Scanner(segments[i]);
            while(currSc.hasNextDouble()) {
                double destKey = currSc.nextDouble();
                double sourceKey = currSc.nextDouble();
                double range = currSc.nextDouble();
                var r = new R(sourceKey, destKey, range);
                // if(hasOverlap(r, ranges)) {
                // System.out.println("overlap found :( ");
                // return;
                // }
                ranges.add(r);
            }
            rangeMaps.add(ranges);
        }

        // the answer for part 1 was less than max int, so this should be fine yolo
        for(double i = 0; i < 1000000000; i++) {
            if(i%10000 == 0) {
                // System.out.println("i is " + i);
            }
            double curNum = i;
            if(curNum == 46.0) {
                System.out.println();
            }
            for(int j = rangeMaps.size() - 1; j >= 0; j--) {
                for(R range: rangeMaps.get(j)) {
                    if(range.dest <= curNum && curNum < range.dest + range.range) {
                        curNum = (curNum - range.dest) + range.source;
                        break;
                    }
                    // Keep curNum - it maps to itself
                }
            }
            for(Seed seed: seeds) {
                if(seed.start <= curNum && curNum < seed.start + seed.range) {
                    System.out.println("The answer is: " + i);
                    System.out.println("This took " + ((System.nanoTime() - start) / 1000000) + " milliseconds to run");
                    return;
                }
            }

        }
        // System.out.println("No overlaps yay");
        System.out.println("We're gunna need a bigger loop");

    }

    // Used this to confirm no overlaps between any ranges, since there's none we can work backwards yeehaw inverses
    public static boolean hasOverlap(R r, List<R> list) {
        for(R cur: list) {
            var a = (cur.source <= r.source && r.source < cur.source + cur.range);
            var b = (r.source <= cur.source && cur.source + cur.range < r.source + r.range);
            var c = (r.source <= cur.source && cur.source < r.source + r.range);
            var d = (cur.source <= r.source && r.source + r.range < cur.source + cur.range);
            if(a // case A
            || b // Case B
            || c  // case D
            || d// Case C
            ) {
                System.out.println("cur: " + cur);
                System.out.println("r: " + r);
                System.out.printf("A: %b B: %b C: %b D: %b ", a, b, c, d);
                return true;
            }
        }
        return false;
    }

}

class Seed{
    double start;
    double range;
    public Seed(double s, double r) {
        start = s;
        range = r;
    }
}

class R {
    double source;
    double dest;
    double range;

    public R(double l, double r, double range) {
        source=l;
        dest=r;
        this.range = range;
    }

    public String toString() {
        return "source: " + source + " , dest: " + dest + "range: " + range;
    }
}
