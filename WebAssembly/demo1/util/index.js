const fs = require('fs')
const path = require('path')
// const result = []
function genTestJson (num) {
    const filePath = path.resolve(__dirname, 'test.json')
    const ws = fs.createWriteStream(filePath)
    ws.write('[')
    for (let i = 0; i < num; i++) {
        // result.push({
        //     index: i,
        //     id: _guid2(),
        //     num: parseInt(Math.random() * 100)
        // })
            const result = {
                index: i,
                id: _guid2(),
                num: parseInt(Math.random() * 100)
            }
       
        if (i === num -1) {
            ws.write(JSON.stringify(result))
        } else {
            ws.write(JSON.stringify(result)+ ',')

        }
        
    }
    ws.write(']')
    ws.end()
}

function _guid2() {
    function S4() {
        return (((1 + Math.random()) * 0x10000) | 0).toString(16).substring(1);
    }
    return (S4() + S4() + "-" + S4() + "-" + S4() + "-" + S4() + "-" + S4() + S4() + S4());
}

genTestJson(7000000)
console.log('json 文件生成了')

// const filePath = path.resolve(__dirname, 'test.json')
// console.log(filePath)
// fs.writeFile(filePath, JSON.stringify(result), () => {
//     console.log('json 文件生成了')
// })