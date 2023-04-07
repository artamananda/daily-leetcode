public class Main{
    public static void main(String[] args) {
        int num = 38;
        System.out.println(addDigits(num));
    }

    public static int addDigits(int num) {
        int result  = 0;
        String str = Integer.toString(num);
        for(int i =0; i < str.length(); i++){
            result += Character.getNumericValue(str.charAt(i));
        }
        while(Integer.toString(result).length() > 1){
            result = addDigits(result);
        }
        return result;
    }
}