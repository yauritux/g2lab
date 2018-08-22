package com.yauritux.service.command;

import java.math.BigDecimal;

import com.yauritux.exception.CardException;
import com.yauritux.model.entity.Card;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 *
 */
public interface CardCommandService {

	Card initializeCard(String owner, BigDecimal initialBalance) throws CardException;
	BigDecimal deduct(String serialNo, BigDecimal amount) throws CardException;
	BigDecimal add(String serialNo, BigDecimal amount) throws CardException;
}
