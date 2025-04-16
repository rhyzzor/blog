const formatDate = (date: Date | string | number, withTime = false): string => {
	const validDate = date instanceof Date ? date : new Date(date);

	if (Number.isNaN(validDate.getTime())) {
		throw new Error("Invalid date provided");
	}

	const options: Intl.DateTimeFormatOptions = {
		year: "numeric",
		month: "2-digit",
		day: "2-digit",
		...(withTime && {
			hour: "2-digit",
			minute: "2-digit",
			second: "2-digit",
			hour12: false,
		}),
	};

	return new Intl.DateTimeFormat("pt-BR", options).format(validDate);
};

export default formatDate;
