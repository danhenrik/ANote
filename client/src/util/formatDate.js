const formatDate = (dateString) => {
  const date = new Date(dateString);

  const dayOptions = {
    day: "2-digit",
    month: "2-digit",
    year: "numeric",
  };

  const hourOptions = {
    hour: "2-digit",
    minute: "2-digit",
  };

  const formattedDay = date.toLocaleDateString(undefined, dayOptions);
  const formattedHour = date.toLocaleTimeString(undefined, hourOptions);

  return { day: formattedDay, hour: formattedHour };
};

export default formatDate;
