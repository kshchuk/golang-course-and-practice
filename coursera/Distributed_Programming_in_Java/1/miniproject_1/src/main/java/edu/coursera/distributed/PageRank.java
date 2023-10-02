package edu.coursera.distributed;

import org.apache.spark.api.java.JavaRDD;

import java.util.LinkedList;
import java.util.List;

import org.apache.spark.api.java.JavaPairRDD;
import scala.Tuple2;

/**
 * A wrapper class for the implementation of a single iteration of the iterative
 * PageRank algorithm.
 */
public final class PageRank {
    /**
     * Default constructor.
     */
    private PageRank() {
    }

    /**
     *
     * @param sites The connectivity of the website graph, keyed on unique
     *              website IDs.
     * @param ranks The current ranks of each website, keyed on unique website
     *              IDs.
     * @return The new ranks of the websites graph, using the PageRank
     *         algorithm to update site ranks.
     */
    public static JavaPairRDD<Integer, Double> sparkPageRank(
            final JavaPairRDD<Integer, Website> sites,
            final JavaPairRDD<Integer, Double> ranks) {
        JavaPairRDD<Integer, Double> newRanks = sites.join(ranks)
                .flatMapToPair(kv -> {
                    Website edges = kv._2()._1();
                    Double currRank = kv._2()._2();
                    List<Tuple2<Integer, Double>> contribs = new LinkedList<>();
                    Iterator<Integer> iter = edges.edgeIterator();
                    while (iter.hasNext())
                        contribs.add(new Tuple2<>(iter.next(), currRank / edges.getNEdges()));
                    return contribs;
                });
        return newRanks.reduceByKey((Double r1, Double r2) -> r1 + r2).mapValues(v -> 0.15 + 0.85 * v);
    }
}
