package com.example.algoexpert

import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.Assertions.assertThrows
import org.junit.jupiter.api.Test

class MinMaxStackTests {

    @Test
    internal fun `it supports stack operations (push and pop)`() {
        val stack = MinMaxStack()

        stack.push(1)
        stack.push(2)
        stack.push(3)

        assertThat(stack.pop()).isEqualTo(3)
    }

    @Test
    internal fun `it supports stack operations (peek)`() {
        val stack = MinMaxStack()

        stack.push(1)

        assertThat(stack.peek()).isEqualTo(1)
    }

    @Test
    internal fun `it supports constant time maximum determination`() {
        val stack = MinMaxStack()

        stack.push(1)

        assertThat(stack.getMax()).isEqualTo(1)

        stack.push(2)
        assertThat(stack.getMax()).isEqualTo(2)

        stack.push(3)
        assertThat(stack.getMax()).isEqualTo(3)
    }

    @Test
    internal fun `it supports constant time minimum determination`() {
        val stack = MinMaxStack()

        stack.push(1)

        assertThat(stack.getMin()).isEqualTo(1)

        stack.push(2)
        assertThat(stack.getMin()).isEqualTo(1)

        stack.push(-33)
        assertThat(stack.getMin()).isEqualTo(-33)
    }

    @Test
    internal fun `it throws IllegalAccessExceptions for invalid operations`() {
        val stack = MinMaxStack()

        assertThrows(IllegalAccessException::class.java) {
            stack.pop()
        }
    }
}