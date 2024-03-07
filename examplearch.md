1. API Gateway:
   The user submits the expression "2+3" through the API Gateway.

2. Https Service:
   The Https Service receives the expression and creates a unique request ID for it. The request ID is returned to the user for future reference.

3. MySQL DB:
   The Https Service stores the expression and the request ID in the MySQL database.

4. SQS Queue:
   The Https Service also adds the expression and the request ID to the SQS Queue.

5. Parser (triggered by SQS message):
   The Parser is triggered by the new message in the SQS Queue. It parses the expression "2+3" and identifies the operation (addition) and the operands (2 and 3).

6. PBFT Client:
   The Parser sends a request to the PBFT Client to perform the addition operation.

7. PBFT Network (nodes):
   The PBFT Client communicates with the PBFT nodes to reach consensus on the result of the addition operation. The nodes exchange messages and vote on the correct result.

8. gRPC Dispatcher:
   Once consensus is reached, the PBFT network communicates the addition operation to the gRPC Dispatcher.

9. Load Balancer (gRPC Add):
   The gRPC Dispatcher forwards the request to the appropriate Load Balancer, which, in this case, is the Load Balancer for the gRPC Add servers.

10. gRPC Add Server:
    The Load Balancer routes the request to one of the gRPC Add servers, which performs the addition operation (2+3) and returns the result (5).

11. PBFT Network (nodes):
    The result (5) is sent back to the PBFT network, which confirms consensus on the result.

12. PBFT Client:
    The PBFT Client receives the consensus result (5) from the PBFT network.

13. Parser:
    The PBFT Client sends the result back to the Parser.

14. MySQL DB:
    The Parser updates the MySQL database with the result (5) corresponding to the request ID.

15. Https Service:
    The result is now stored in the MySQL database and can be retrieved by the Https Service when the user queries the result using the request ID.

16. API Gateway:
    When the user requests the result using the request ID, the API Gateway forwards the request to the Https Service, which retrieves the result (5) from the MySQL database and returns it to the user.
