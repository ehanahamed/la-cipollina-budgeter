export function plainTextTableFrom2DArray(array, padding) {
    let rows = array;
    const colWidths = rows[0].map((_, colIndex) =>
        Math.max(...rows.map(row => row[colIndex].length))
    );
    
    let result = ""
    rows.forEach(row => {
        result += row.map(
            (cell, i) => cell.padEnd(colWidths[i] + padding)
        ).join('') + "\n";
    });
    return result;
}
