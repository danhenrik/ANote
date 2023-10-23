import React, { useState } from "react";
import FavoriteIcon from "@mui/icons-material/Favorite";

const LikeButton = (countLikes) => {
  const [likes, setLikes] = useState(countLikes.countLikes);
  const [isClicked, setIsClicked] = useState(false);

  const handleClick = () => {
    if (isClicked) {
      setLikes(likes - 1);
    } else {
      setLikes(likes + 1);
    }
    setIsClicked(!isClicked);
  };

  return (
    <span>
      <FavoriteIcon
        className={`like-button ${isClicked && "liked"}`}
        onClick={handleClick}
        style={{ marginTop: "10px", color: "red" }}
      />
      <span
        style={{
          color: "red",
          position: "relative",
          bottom: "7px",
          left: "2px",
        }}
        className='likes-counter'
      >
        {`${likes}`}
      </span>
    </span>
  );
};

export default LikeButton;
