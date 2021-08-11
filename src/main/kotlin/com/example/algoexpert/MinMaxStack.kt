package com.example.algoexpert

class MinMaxStack {
    private interface Entry {
        fun min(): Int
        fun max(): Int
        fun value(): Int
    }

    private val entries = mutableListOf<Entry>();

    fun push(value: Int) {
        if (entries.isEmpty()) {
            val entry = object : Entry {
                override fun min() = value
                override fun max() = value
                override fun value() = value
            }

            entries.add(entry);
            return
        }

        val top = entries[entries.size-1]

        val min = if (value < top.min()) value else top.min()
        val max = if (value > top.max()) value else top.max()

        val entry = object : Entry {
            override fun min() = min
            override fun max() = max
            override fun value() = value
        }

        entries.add(entry);
    }

    fun pop(): Int {
        if (entries.isEmpty()) {
            throw IllegalAccessException("no such elements")
        }

        return entries.removeAt(entries.size-1).value()
    }

    fun peek(): Int {
        if (entries.isEmpty()) {
            throw IllegalAccessException("no such elements")
        }

        return entries[entries.size-1].value()
    }

    fun getMax(): Int {
        if (entries.isEmpty()) {
            throw IllegalAccessException("no such elements")
        }

        return entries[entries.size-1].max()
    }

    fun getMin(): Int {
        if (entries.isEmpty()) {
            throw IllegalAccessException("no such elements")
        }

        return entries[entries.size-1].min()
    }

}
