/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/
package org.apache.plc4x.codegen.ast;

import java.util.Arrays;
import java.util.List;

public class Block extends Statement {

    public static final Block EMPTY_BLOCK = new Block();

    private final List<Node> statements;

    Block(List<Node> statements) {
        this.statements = statements;
    }

    Block(Node... statements) {
        this(Arrays.asList(statements));
    }

    public List<Node> getStatements() {
        return statements;
    }

    public static BlockBuilder build() {
        return new BlockBuilder();
    }

    @Override public <T> T accept(NodeVisitor<T> visitor) {
        return null;
    }

    @Override public void write(Generator writer) {
        writer.writeBlock(this);
    }
}
