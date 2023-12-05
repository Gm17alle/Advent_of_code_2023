import java.io.File;
import java.util.Scanner;
import java.io.FileNotFoundException;

/**
 * Write a description of class Part1 here.
 *
 * @author (your name)
 * @version (a version number or a date)
 */
public class Part1
{
    public static void main(String[] args) {
        char[][] grid = null;
        try {
            grid = readFile();
        } catch (Throwable t) {
            //do nothing
        }
        System.out.println("Grid length: " + grid.length);
        System.out.println("Grid Width: " + grid[0].length);
        int sum = 0;
        for(int i = 0; i < grid.length; i++) {
            for(int j = 0; j < grid[0].length; j++) {
                String newNum = "";
                int jStart = j;
                while(j < grid.length && isIntChar(grid[i][j])) {
                    newNum += grid[i][j];
                    j++;
                }
                if(newNum.length() > 0 && isValid(grid, jStart, j, i)) {
                    sum += Integer.parseInt(newNum);
                } else if(newNum.length() > 0) {
                    System.out.println("Ignoring num: " + newNum);
                }
            }
        }
        System.out.println(sum);

    }
    
    public static boolean isIntChar(char c) {
        return c - '0' >= 0 && c - '0' <= 9;
    }

    public static boolean isValid(char[][] grid, int jStart, int j, int i) {
        //TODO
        for(int start = i-1; start <= i+1; start++) {
            for(int end = jStart-1; end <= j; end++) {
                if(start > -1 && start < grid.length && end > -1 && end < grid[0].length
                && !isIntChar(grid[start][end]) && grid[start][end] != '.') {
                    return true;
                }
            }
        }
        return false;   
    }

    public static char[][] readFile() throws FileNotFoundException {
        char[][] grid = new char[140][140];
        File f = new File("input.txt");
        Scanner sc = new Scanner(f);

        int i = 0;
        while(sc.hasNext()) {
            String line = sc.nextLine();
            grid[i] = line.toCharArray();
            i++;
        }

        return grid;
    }
}
