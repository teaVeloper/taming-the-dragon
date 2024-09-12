import XCTest
import SwiftTreeSitter
import TreeSitterDragonc

final class TreeSitterDragoncTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_dragonc())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Dragonc grammar")
    }
}
