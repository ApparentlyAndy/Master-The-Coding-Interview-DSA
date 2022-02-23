class HashTable {
    constructor(size) {
        this.data = new Array(size);
    }

    _hash(key) {
        let hash = 0;
        for (let i = 0; i < key.length; i++) {
            hash = (hash + key.charCodeAt(i) * i) % this.data.length
        }
        return hash;
    }

    set(key, data) {
        const hash = this._hash(key)
        if (!this.data[hash]) {
            this.data[hash] = [[key, data]]
        } else {
            this.data[hash].push([key, data])
        }
        console.log(this.data)
    }

    get(key) {
        const hash = this._hash(key)
        if (this.data[hash].length > 1) {
            this.data[hash].forEach(entry => {
                if (hash === entry[0]) {
                    return entry[1]
                }
            })
        } else {
            console.log(this.data[hash][0][1])
            return this.data[hash][0][1]
        }
        console.log('not found')
        return undefined
    }
}

const myHashTable = new HashTable(2);
myHashTable.set('grapes', 10000)
myHashTable.get('grapes')
myHashTable.set('apples', 9)
myHashTable.get('apples')