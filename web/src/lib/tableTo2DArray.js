export function tableTo2DArray(table) {
    let result = [];
    for (
        let rowIndex = 0;
        rowIndex < table.rows.length;
        rowIndex++
    ) {
        let row = [];
        for (
            let cellIndex = 0;
            cellIndex < table.rows[rowIndex].cells.length;
            cellIndex ++
        ) {
            row.push(
                table.rows[rowIndex].cells[
                    cellIndex
                ].innerText
            );
        }
        result.push(row);
    }
    return result;
}
