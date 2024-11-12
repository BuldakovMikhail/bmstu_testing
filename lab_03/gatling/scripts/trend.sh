#!/bin/bash

rm -rf serverpersecondloadsimulation-trend
java -jar ~/gatling-report-6.1-capsule-fat.jar $(ls /mnt/c/sem7/bmstu_testing/lab_03/gatling/results/serverpersecondloadsimulation*/simulation.log) -o serverpersecondloadsimulation-trend

rm -rf serveratonceloadsimulation-trend
java -jar ~/gatling-report-6.1-capsule-fat.jar $(ls /mnt/c/sem7/bmstu_testing/lab_03/gatling/results/serveratonceloadsimulation*/simulation.log) -o serveratonceloadsimulation-trend