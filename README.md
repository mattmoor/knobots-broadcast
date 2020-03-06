# knobots broadcaster

The purpose of this repository is to expose the Knative Github events to the broader community to enable more folks to write bots that enable workflow improvements in the same way that I have been working on the [knobots](https://github.com/mattmoor/knobots).

There are three key parts:
1. The sources (`config/hooks`) which subscribe to all of the Github event data for a collection of the Knative repositories.
2. The triggers (`config/triggers`) which folks can open PRs to receive our Github event stream.
3. The "receiver" (`config/receiver`, `cmd/receiver`) which you run in your own cluster in order to receive events.

## Previewing the event stream

You can get a preview of the event stream flowing from Knative using Scotty's [sockeye](https://github.com/n3wscott/sockeye) tool [here](http://display.default.broadcast.knobots.io).

## Receiving events

In order to receive your own events, configure a cluster with Knative's [Serving and Eventing components](https://knative.dev/docs/install/any-kubernetes-cluster/), and enable the Broker on the namespace where you would
like to receive events.  Then run:

```shell
ko apply -n default -f config/receiver
```

Watch for the `receiver` `ksvc` to become Ready and extract it's URL.


Next, open a PR against this repository adding your own trigger, here is the one powering the knobots:

```yaml
apiVersion: eventing.knative.dev/v1alpha1
kind: Trigger
metadata:
  name: knobots-trigger         <-- Give this a unique name
  namespace: default
spec:
  filter: {}
  subscriber:
    uri: http://receiver.default.knobots.io   <-- Change to your URL
```

Once you have opened the PR, ping me on [slack.knative.dev](https://slack.knative.dev).


## Disclaimer & Warning

This is not a product, it has no SLA, and if you abuse it, I will remove your trigger.
