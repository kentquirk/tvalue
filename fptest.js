
function threshold(tValue) {
	const k = 0x100_0000_0000_0000  // 2^56

	if (tValue < 1.0) {
		return Math.floor(k*tValue + 0.5)
	}
	return Math.floor(k/tValue + 0.5)
}

function samplingRate(tValue) {
	if (tValue < 1.0) {
		return Math.floor(1.0/tValue + 0.5)
	}
	return Math.floor(tValue + 0.5)
}

function samplingProbability(tValue) {
	if (tValue < 1.0) {
		return tValue
	}
	return 1.0 / tValue
}

function all(args) {
    for (var i=0; i<args.length; i++) {
        var arg = process.argv[i]
        var tValue = parseFloat(arg)
        if (isNaN(tValue)) {
            console.log("bad value: %s", arg)
            continue
        }

		console.log("tValue: %f", tValue)
		console.log("threshold: %d", threshold(tValue))
		console.log("samplingRate: %d", samplingRate(tValue))
		console.log("samplingProbability: %f", samplingProbability(tValue))
		console.log()
	}
}

function thresh(arg) {
	var tValue = parseFloat(arg)
	if (isNaN(tValue)) {
		return "bad value: " + arg
	}
	return ""+threshold(tValue)
}

function rate(arg) {
	var tValue = parseFloat(arg)
	if (isNaN(tValue)) {
		return "bad value: " + arg
	}
	return ""+samplingRate(tValue)
}

function prob(arg) {
	var tValue = parseFloat(arg)
	if (isNaN(tValue)) {
		return "bad value: " + arg
	}
	return samplingProbability(tValue).toFixed(14)
}

function main() {
	var args = process.argv.slice(2)
	if (args.length == 0) {
		console.log("usage: node fptest.js command <tValue> ...")
		return
	}
	if (args[0] == "all") {
		all(args)
	} else if (args[0] == "thresh") {
		console.log(thresh(args[1]))
	} else if (args[0] == "rate") {
		console.log(rate(args[1]))
	}
	else if (args[0] == "prob") {
		console.log(prob(args[1]))
	} else {
		console.log("bad command: %s", args[0])
	}
}

main()

