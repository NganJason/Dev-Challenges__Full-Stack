import React, { useState, useEffect } from "react"
import Upload from "../Upload/Upload";
import Uploading from "../Uploading/Uploading";
import Uploaded from "../Uploaded/Uploaded";

function Dashboard() {
  const [file, setFile] = useState(null)
  const [submitted, setSubmitted] = useState(false)

  const onFileSubmit = (file) => {
    setFile(file);

    setInterval(() => {
      setFile(null)
      setSubmitted(true)
    }, 500)
  };

  const getComponent= () => {
    if (file === null && !submitted) {
      return <Upload onFileSubmit={onFileSubmit} />;
    } else if (submitted) {
      return <Uploaded/>
    }

    return <Uploading/>
  }

  return <div className="dashboard">
  {
    getComponent()  
  }
  </div>;
}

export default Dashboard;
