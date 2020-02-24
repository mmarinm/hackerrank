function minimumSwaps(arr) {
    let ans = 0;
    const visited = {};
    const startState = arr.map((val, idx) => {
        visited[idx] = false;
        return [idx, val];
    });

    // sort based on arr values
    const endState = startState.sort((a, b) => a[1] - b[1]);

    for (let i = 0; i < arr.length; i++) {
        if (visited[i] || endState[i][0] === i) {
            continue;
        }

        answ += calcCycleSize(i, endState, visited);
    }
    return ans;
}

//traverse graph
function calcCycleSize(i, arr, visited) {
    let cycleSize = 0;
    let j = i;
    // while not full circle
    while (!visited[j]) {
        visited[j] = true;
        j = arr[j][0];
        cycleSize++;
    }
    if (cycleSize > 0) {
        return cycleSize - 1;
    }
    return 0;
}
