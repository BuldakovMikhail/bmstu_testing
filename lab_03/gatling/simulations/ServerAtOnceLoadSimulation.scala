import io.gatling.core.Predef._
import io.gatling.http.Predef._
import scala.concurrent.duration._

class ServerAtOnceLoadSimulation extends Simulation {

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
        atOnceUsers(50000),
        nothingFor(5.seconds)
    ).protocols(httpProtocolEcho),
    scnFastHTTP.inject(
        atOnceUsers(50000),
        nothingFor(5.seconds)
    ).protocols(httpProtocolFastHTTP),
    scnChi.inject(
        atOnceUsers(50000),
        nothingFor(5.seconds)
    ).protocols(httpProtocolChi)
  ).maxDuration(60.seconds)
}