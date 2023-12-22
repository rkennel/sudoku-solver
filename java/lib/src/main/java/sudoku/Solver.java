package sudoku;

public class Solver {
    private Solver(){
        super();
    }

    public static final int VALUE_UNKNOWN=0;

    /**
     * Solves an unfinished Sudoku Puzzle
     * @param boardArray a multidimensional integer array
     *                   first position in array represents row
     *                   second position in array represents column
     *                   if the value is unknown at a given position in array then it is inputted with a value of 0
     * @return a multidimensional array with the final solution following the same conventions as the input array
     */
    public static int[][] solve(int[][] boardArray) {
        return new int[9][9];
    }
}
