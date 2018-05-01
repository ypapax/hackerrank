package com.mycompany.app;

import org.apache.log4j.Logger;

import java.util.Arrays;
import java.util.Scanner;

public class App 
{
    static Logger logger = Logger.getLogger(App.class);

    public static void main(String[] args) {
        System.out.println("getLocation");
        System.out.println(App.class.getProtectionDomain().getCodeSource().getLocation());
        logger.info("INFO2");
        Scanner sc = new Scanner(System.in);

        int q = sc.nextInt();
        System.out.printf("q %+d\n", q);
        for (int tc = 0; tc < q; tc++) {
            int n = sc.nextInt();
            int[][] M = new int[n][n];
            for (int i = 0; i < n; i++) {
                for (int j = 0; j < n; j++) {
                    M[i][j] = sc.nextInt();
                }
            }

            System.out.println(solve(M) ? "Possible" : "Impossible");
        }

        sc.close();
    }

    static boolean solve(int[][] M) {
        int n = M.length;

        long[] rowSums = new long[n];
        for (int r = 0; r < n; r++) {
            for (int c = 0; c < n; c++) {
                rowSums[r] += M[r][c];
            }
        }

        long[] colSums = new long[n];
        for (int c = 0; c < n; c++) {
            for (int r = 0; r < n; r++) {
                colSums[c] += M[r][c];
            }
        }

        return isSame(rowSums, colSums);
    }

    static boolean isSame(long[] a, long[] b) {
        Arrays.sort(a);
        Arrays.sort(b);

        for (int i = 0; i < a.length; i++) {
            if (a[i] != b[i]) {
                return false;
            }
        }
        return true;
    }
}
