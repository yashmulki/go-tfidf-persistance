# go-tfidf-persistance
(IN PROGRESS) TFIDF in Go with support for persisting document frequency data

TFIDF is a useful measure for the salience of words in a document to the overall meaning of a document. It uses a formula that leverages information about the frequency of words both within a singular document as well as in a corpus of documents. It is widely applied in NLP tasks.

This library allows you to easily create and manage a corpus of documents for TFIDF and save it in a json format. This makes it very easy to perform TFIDF while simulataneously building up the overall document corpus, and quickly save/load this data to disk. 
