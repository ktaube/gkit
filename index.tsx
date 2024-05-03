#! /usr/bin/env bun

import { $ } from "bun";
import React, { useState, useEffect } from "react";
import { render, Text } from "ink";

const { stderr, stdout } = await $`git --version`.nothrow().quiet();

console.log("stdout", stdout.toString());
console.log("stderr", stderr.toString());

const Counter = () => {
  const [counter, setCounter] = useState(0);

  useEffect(() => {
    const timer = setInterval(() => {
      setCounter((previousCounter) => previousCounter + 1);
    }, 100);

    return () => {
      clearInterval(timer);
    };
  }, []);

  return <Text color="green">{counter} tests passed</Text>;
};

render(<Counter />);
