# sharego

An experimental port of sharejs to golang. Not recommended for the real world.

Implements the server half of an operational transform session, allowing multiple
users to collaborate on a document in real time.

At this stage there's no network support, that's to come. The intention is to
offer REST and browserchannel interfaces.

## Installation

    go get github.com/yob/sharego

## Usage

At this stage, there's no way to interactively use sharego. There's a test app
you can run to see document editing in action.

    go install github.com/yob/sharego/sharego-play
    ./bin/sharego-play
