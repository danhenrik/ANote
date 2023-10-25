import React, { useState } from "react";
import FavoriteIcon from "@mui/icons-material/Favorite";
import FavoriteBorderIcon from "@mui/icons-material/FavoriteBorder";
import axios from "axios";

const LikeButton = (note) => {
  const [likes, setLikes] = useState(note.note.LikesCount);
  const [isClicked, setIsClicked] = useState(false);

  const handleClick = async () => {
    try {
      if (isClicked) {
        axios.delete("/likes/" + note.note.AuthorId + "/" + note.note.Id);
        setLikes(likes - 1);
      } else {
        const likeData = {
          user_id: note.note.AuthorId,
          note_id: note.note.Id,
        };
        axios.post("/likes", likeData);
        setLikes(likes + 1);
      }
      setIsClicked(!isClicked);
    } catch (error) {
      console.log("Like failed:", error);
    }
  };

  return (
    <span>
      {isClicked ? (
        <FavoriteIcon
          className={`like-button ${isClicked && "liked"}`}
          onClick={handleClick}
          style={{ marginTop: "10px", color: "red" }}
        />
      ) : (
        <FavoriteBorderIcon
          className={`like-button ${isClicked && "liked"}`}
          onClick={handleClick}
          style={{ marginTop: "10px", color: "red" }}
        />
      )}
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
