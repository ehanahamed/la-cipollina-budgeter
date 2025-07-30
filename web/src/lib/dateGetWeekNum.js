export function dateGetWeekNum(date) {
    const jan01 = new Date(date.getFullYear(), 0, 1);
    const daysSinceJan01 = Math.floor((date - jan01) / 86400000);
    const firstWeekday = jan01.getDay();
    const offset = (firstWeekday === 0) ?
        6 : firstWeekday - 1;
    return Math.floor((daysSinceJan01 + offset) / 7) + 1;
}
