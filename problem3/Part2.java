import java.util.Map;
import java.util.HashMap;
import java.util.List;
import javafx.util.Pair;
import java.util.LinkedList;
import java.util.Collection;

/**
 * Write a description of class Part2 here.
 *
 * @author (your name)
 * @version (a version number or a date)
 */
public class Part2
{
    public static void main(String args[]) {
        char[][] grid = null;
        try {
            grid = Part1.readFile();
        } catch (Throwable t) {
            //do nothing
        }
        var map = new HashMap<Pair<Integer, Integer>, List<Integer>>(); 
        for(int i = 0; i < grid.length; i++) {
            for(int j = 0; j < grid[0].length; j++) {
                String newNum = "";
                int jStart = j;
                while(j < grid.length && Part1.isIntChar(grid[i][j])) {
                    newNum += grid[i][j];
                    j++;
                }
                if(newNum.length() > 0) {                    
                    handleIt(grid, Integer.parseInt(newNum), jStart, j, i, map);
                } else if(newNum.length() > 0) {
                    System.out.println("Ignoring num: " + newNum);
                }
            }
        }
        int sum = 0;
        Collection<List<Integer>> values = map.values();
        for(List<Integer> v: values) {
            if(v.size() == 2) {
                sum += v.get(0) * v.get(1);
            }
        }
        System.out.println(sum);
    }
    
    public static void handleIt(char[][] grid, int newNum, int jStart, int j, int i, Map<Pair<Integer, Integer>, List<Integer>> map) {
        //TODO
        for(int start = i-1; start <= i+1; start++) {
            for(int end = jStart-1; end <= j; end++) {
                if(start > -1 && start < grid.length && end > -1 && end < grid[0].length
                && grid[start][end] == '*') {
                    Pair<Integer, Integer> p = new Pair<Integer, Integer>(start, end);
                    if(!map.containsKey(p)) {
                        List<Integer> list = new LinkedList<Integer>();
                        map.put(p, list);
                    }
                    List<Integer> list = map.get(p);
                    list.add(newNum);
                }
            }
        }  
    }
    
    
}
