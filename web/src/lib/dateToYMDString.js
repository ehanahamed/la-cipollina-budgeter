export function dateToYMDString(dateObj) {
    return dateObj.getFullYear() + "-" +
        (dateObj.getMonth() + 1).toString().padStart(2, "0") + "-" +
        dateObj.getDate().toString().padStart(2, "0")
}
