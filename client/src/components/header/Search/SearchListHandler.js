const listHandler = (setList) => {
  const addToList = (item) => {
    setList((prevList) => {
      if (!prevList.includes(item)) {
        return [...prevList, item];
      } else {
        return prevList;
      }
    });
  };

  const removeFromList = (item) => {
    setList((prevList) => prevList.filter((listItem) => listItem !== item));
  };

  return {
    addToList,
    removeFromList,
  };
};

export default listHandler;
