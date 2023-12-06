import java.math.BigInteger;

/**
 * Write a description of class hello here.
 *
 * @author (your name)
 * @version (a version number or a date)
 */
public class whyisthisday6thisistheeasiestonethisisbonkerslol
{
    public static void main(String args[]) {
        int[] times = {56, 97, 78, 75};
        var time = new BigInteger("56977875");
        var distance = new BigInteger("546192711311139");
        int []distances = {546, 1927, 1131, 1139};
        int product = 1;
        long count = 0;
        var one = new BigInteger("1");
        var million = new BigInteger("1000000");
        var j = new BigInteger("1");
        while(j.compareTo(time) == -1) {
            if(time.subtract(j).multiply(j).compareTo(distance) > 0) {
                count++;
            }
            j = j.add(one);
        }

        System.out.println(count);

    }
}
