import { createContext, useContext, useState } from "react";
import PropTypes from "prop-types";
import ModalComponent from "../common/ModalComponent";

const ModalContext = createContext();

export const ModalProvider = ({ children }) => {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [modalContent, setModalContent] = useState(null);

  const defaultStyle = {
    display: "flex",
    alignItems: "flex-start",
    justifyContent: "center",
    overflow: "auto",
    maxHeight: "100vh",
    "@media (min-height: 400px)": {
      alignItems: "center",
    },
  };

  const [modalStyling, setModalStyling] = useState(defaultStyle);

  const openModal = (content) => {
    setModalContent(content);
    setIsModalOpen(true);
  };

  const closeModal = () => {
    setIsModalOpen(false);
    setModalContent(null);
  };

  return (
    <ModalContext.Provider value={{ openModal, closeModal, setModalStyling }}>
      {children}
      <ModalComponent
        content={modalContent}
        open={isModalOpen}
        handleClose={closeModal}
        style={modalStyling}
      />
    </ModalContext.Provider>
  );
};

export function useModal() {
  return useContext(ModalContext);
}

ModalProvider.propTypes = {
  children: PropTypes.node.isRequired,
};
