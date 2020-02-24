// BruteForce O(n^2)
function minimumBribesBF(q) {
    let count = 0;
    for (let i = 0; i < q.length; i++) {
        let countIter = 0;
        const el = q[i];
        for (let j = q.length - 1; j > i; j--) {
            if (el > q[j]) {
                count++;
                countIter++;
                if (countIter > 2) {
                    console.log("Too chaotic");
                    return;
                }
            }
        }
    }
    console.log(count);
}

const swap = (p1, p2, arr) => {
    const temp = arr[p2]; //8
    arr[p2] = arr[p1];
    arr[p1] = temp;
    return arr;
};

function minimumBribes(q) {
    const n = q.length; //7
    let cnt = 0;
    for (let i = n - 1; i >= 0; i--) {
        if (q[i] !== i + 1) {
            //some swaping has happened
            //swap once
            if (i - 1 >= 0 && q[i - 1] === i + 1) {
                cnt++;
                q = swap(i, i - 1, q);
            }

            //swap twice
            else if (i - 2 >= 0 && q[i - 2] === i + 1) {
                cnt += 2;
                q = swap(i - 1, i - 2, q);
                q = swap(i, i - 1, q);
            } else {
                //swap too many times
                console.log("Too chaotic");
                return;
            }

        }
    }
    console.log(cnt);
    return
}

minimumBribes([2, 1, 5, 3, 4]);
minimumBribes([2, 5, 1, 3, 4]);
minimumBribes([2, 5, 1, 3, 4]);
minimumBribes([5, 1, 2, 3, 7, 8, 6, 4]);
minimumBribes([1, 2, 5, 3, 4, 7, 8, 6]);
