export default (date: Date) => {
 const newDate = new Date(date).toISOString().split('T')[0];

 return newDate;
}
