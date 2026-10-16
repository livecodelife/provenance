import os
import sys
import unittest

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
import metrics  # noqa: E402


class MetricsTest(unittest.TestCase):
    def test_recall_single_truth(self):
        self.assertEqual(metrics.recall_at(["a", "b", "c"], ["b"], 1), 0.0)
        self.assertEqual(metrics.recall_at(["a", "b", "c"], ["b"], 3), 1.0)

    def test_recall_caps_at_k(self):
        # Two true records and k=1: finding one of them is full marks.
        self.assertEqual(metrics.recall_at(["a", "x"], ["a", "b"], 1), 1.0)
        self.assertEqual(metrics.recall_at(["a", "x", "y"], ["a", "b"], 3), 0.5)

    def test_reciprocal_rank(self):
        self.assertEqual(metrics.reciprocal_rank(["x", "y", "b"], ["b"]), 1 / 3)
        self.assertEqual(metrics.reciprocal_rank(["x"], ["b"]), 0.0)
        self.assertEqual(metrics.reciprocal_rank(["x"] * 10 + ["b"], ["b"]), 0.0)

    def test_rrf_prefers_agreement(self):
        fused = metrics.rrf([["a", "b", "c"], ["d", "b", "c"]])
        self.assertEqual(fused[0], "b")
        self.assertEqual(set(fused), {"a", "b", "c", "d"})

    def test_summarize(self):
        s = metrics.summarize([(["a"], ["a"]), (["x", "a"], ["a"])])
        self.assertEqual(s["n"], 2)
        self.assertEqual(s["R@1"], 0.5)
        self.assertEqual(s["R@3"], 1.0)
        self.assertEqual(s["MRR"], 0.75)
        self.assertEqual(metrics.summarize([])["n"], 0)


if __name__ == "__main__":
    unittest.main()
