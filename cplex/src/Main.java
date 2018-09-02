import domain.Thing;
import ilog.concert.IloException;
import ilog.concert.IloLinearNumExpr;
import ilog.concert.IloNumVar;
import ilog.cplex.IloCplex;

public class Main {
    public static void main(String[] args) {

        final Thing[] things = {
                new Thing(10, 10),
                new Thing(100, 10),
        };
        final int T = things.length;

        final int capacity = 19;

        final double mu = 10.0;

        final int iterations = 5;

        double[] xLast = new double[T];

        for (int iteration = 0; iteration < iterations; iteration++) {
            try {
                IloCplex cplex = new IloCplex();

                // binary variable assuming the value 1 if the _h_th thing is stored on knapsack;
                // otherwise its value is zero
                String[] xNames = new String[T];
                for (int i = 0; i < T; i++) {
                    xNames[i] = String.format("x(%d)", i);
                }
                IloNumVar[] x = cplex.numVarArray(T, 0, 1, xNames);

                // objective function
                IloLinearNumExpr expr = cplex.linearNumExpr();
                for (int i = 0; i < T; i++) {
                    expr.addTerm(things[i].getCost(), x[i]);
                    expr.addTerm(mu * (2 * xLast[i] - 1), x[i]);
                }
                cplex.addMaximize(expr);

                // weight constraint
                IloLinearNumExpr constraint = cplex.linearNumExpr();
                for (int i = 0; i < T; i++) {
                    constraint.addTerm(things[i].getWeight(), x[i]);
                }
                cplex.addLe(constraint, capacity);

                if (cplex.solve()) {
                    System.out.println();
                    System.out.println(" Solution Status = " + cplex.getStatus());
                    System.out.println();
                    System.out.println(" cost = " + cplex.getObjValue());

                    System.out.println();
                    System.out.println(" >> Things");
                    for (int i = 0; i < T; i++) {
                        System.out.printf("x[%d] = %g\n", i, cplex.getValue(x[i]));
                        xLast[i] = cplex.getValue(x[i]);
                    }
                    System.out.println();

                } else {
                    System.out.printf("Solve failed: %s\n", cplex.getStatus());
                }

            } catch (IloException e) {
                e.printStackTrace();
            }
        }
    }
}
