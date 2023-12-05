import java.util.HashSet;
import java.util.Set;
import java.io.File;
import java.util.Scanner;
import java.io.FileNotFoundException;
import java.util.HashMap;
import java.util.Map;
import java.util.Collections;
import java.util.*;

/**
 * Write a description of class MapCity here.
 *
 * @author (your name)
 * @version (a version number or a date)
 */
public class MapCity
{
    public static void main(String args[]) throws FileNotFoundException { 
        Set<Double> curSet = new HashSet<Double>();
        File f = new File("maps.txt");
        Scanner sc = new Scanner(f);
        String s = "";
        while(sc.hasNextLine()) {
            s += sc.nextLine() + "\n";
        }
        var segments = s.split(":");
        Scanner scanSeeds = new Scanner(segments[1]);
        while(scanSeeds.hasNextDouble()) {
            curSet.add(scanSeeds.nextDouble());
        }
        
        for(int i = 2; i < segments.length; i++) {
            var newSet = new HashSet<Double>();
            var mapScanner = new Scanner(segments[i]);
            // var curMap = new HashMap<Double, Double>();
            while(mapScanner.hasNextDouble()) {
                var dest = mapScanner.nextDouble();
                var source = mapScanner.nextDouble();
                var range = mapScanner.nextDouble();
                for(double j = dest; j < dest+range; j++) {
                    // curMap.put(source, j);
                    if(curSet.contains(source)) {
                        curSet.remove(source);
                        newSet.add(j);
                    }
                    source++;
                }
                // System.out.println("Current map: ");
                // curMap.forEach((key, value) -> System.out.println(key + " " + value));
            }
            // System.out.println("Current map: ");
            // curMap.forEach((key, value) -> System.out.println(key + " " + value));
            for(var val: curSet) {
                newSet.add(val);
            }
            curSet = newSet;
            System.out.println("i is now: " + i + "\n Set is now " + Arrays.toString(curSet.toArray()));
        }
        System.out.println(Collections.min(curSet));
    }
}
