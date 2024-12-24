import io.gatling.core.Predef._
import io.gatling.http.Predef._
import scala.concurrent.duration._

class ServerPerSecondLoadSimulation extends Simulation {

  val httpProtocolEcho = http.baseUrl("http://echo-ping:8081")
  val httpProtocolFastHTTP = http.baseUrl("http://fasthttp-ping:8082")
  val httpProtocolChi = http.baseUrl("http://chi-ping:8083")

  val scnEcho = scenario("Echo Server Load Test")
    .exec(http("Echo Metrics").get("/ping"))
  val scnFastHTTP = scenario("FastHTTP Server Load Test")
    .exec(http("FastHTTP Metrics").get("/ping"))
  val scnChi = scenario("Chi Server Load Test")
      .exec(http("Chi Metrics").get("/ping"))

  setUp(
    scnEcho.inject(
        constantUsersPerSec(2000).during(30)
    ).protocols(httpProtocolEcho),
    scnFastHTTP.inject(
        constantUsersPerSec(2000).during(30)
    ).protocols(httpProtocolFastHTTP),
    scnChi.inject(
        constantUsersPerSec(2000).during(30)
    ).protocols(httpProtocolChi)
  ).maxDuration(60.seconds)
}