public class Main{
    public static void main(String[] args) {
        int num = 55;
        System.out.println(isUgly(num));
    }

    public static boolean isUgly(int n) {
        if(n == 0){
            return false;
        }
        else if(n == 1){
            return true;
        }

        while(n % 2 == 0 || n % 3 == 0 || n % 5 == 0){
            if(n % 2 == 0){
                n/=2;
            }
            else if(n % 3 == 0){
                n/=3;
            }
            else{
                n/=5;
            }
        }

        if(n == 1 || n == 2 || n == 3 || n == 5){
            return true;
        }

        return false;
    }
}