// 金額（数値）を円表記に変換する.
export const formatMoney = (money: string | number, showUnit = true): string => {
  return `${showUnit ? '￥' : ''}${money.toLocaleString()}`
}
